package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/utils"
)

type Session struct {
	UserId       int
	Expiry       time.Time
	Token        string
	ActivateCode string
	UserIp       string
	Email        string
	Type         string
}

type SessionStore struct {
	AuthSessions  map[string]*Session
	sMu           sync.RWMutex
	EmailSessions map[string]*Session
	eMu           sync.RWMutex
	data          data.Data
}
type AuthSessions = map[string]*Session
type EmailSessions = map[string]*Session

func NewSessionStore(d data.Data) *SessionStore {
	return &SessionStore{
		AuthSessions:  make(AuthSessions),
		EmailSessions: make(EmailSessions),
		data:          d,
	}
}

const (
	typeEmail = "email"
	typeAuth  = "auth"
)

func (as *SessionStore) NewAuthSession(r *http.Request, uId int) (*Session, string) {
	newToken := uuid.NewString()
	expiresAt := time.Now().Add(12 * time.Hour)

	us := &Session{
		UserId: uId,
		Expiry: expiresAt,
		Token:  newToken,
		UserIp: utils.GetIP(r),
		Type:   typeEmail,
	}
	log.Println("Generated a new UserSession: ", us)
	return us, newToken
}

func (as *SessionStore) NewEmailSession(r *http.Request, email string, uId int) (*Session, string) {
	expiresAt := time.Now().Add(24 * time.Hour)
	newToken := uuid.NewString()

	hash := sha256.New()
	hash.Write([]byte(newToken))
	activateCode := hex.EncodeToString(hash.Sum(nil))
	ea := &Session{
		UserId:       uId,
		Token:        newToken,
		Expiry:       expiresAt,
		ActivateCode: activateCode,
		UserIp:       utils.GetIP(r),
		Email:        email,
		Type:         typeAuth,
	}
	return ea, newToken
}

func (as *SessionStore) GetAuthSession(token string) (*Session, bool) {
	as.sMu.RLock()
	defer as.sMu.RUnlock()
	session, exists := as.AuthSessions[token]
	if !exists {
		return nil, false
	}
	return session, true
}

func (as *SessionStore) AddAuthSession(token string, session *Session) {
	as.sMu.Lock()
	defer as.sMu.Unlock()
	as.AuthSessions[token] = session
}
func (as *SessionStore) RemoveAuthSession(token string) {
	as.sMu.Lock()
	defer as.sMu.Unlock()
	delete(as.AuthSessions, token)
}

func (as *SessionStore) GetEmailSession(token string) (*Session, bool) {
	as.sMu.RLock()
	defer as.sMu.RUnlock()
	session, exists := as.EmailSessions[token]
	log.Println(session)
	if !exists {
		return nil, false
	}
	return session, true
}

func (as *SessionStore) AddEmailSession(token string, session *Session) {
	as.eMu.Lock()
	defer as.eMu.Unlock()
	as.EmailSessions[token] = session
}
func (as *SessionStore) RemoveEmailSession(token string) {
	as.eMu.Lock()
	defer as.eMu.Unlock()
	delete(as.EmailSessions, token)
}

func (as *SessionStore) AuthorizeAdmin(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		sessionToken := cookie.Value
		var s *Session
		s, exists := as.GetAuthSession(sessionToken)
		log.Println(s)
		if !exists {
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}
		userIp := utils.GetIP(r)
		userId := int(s.UserId)

		if s.UserIp != userIp {
			log.Printf("[!!!] Unauthorized ip: %s wanted to authorize as userID: %v ", userIp, userId)
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if s.Expiry.Before(time.Now()) {
			delete(as.AuthSessions, sessionToken)
			http.Redirect(w, r, "/auth/refresh-token", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), &Session{}, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
