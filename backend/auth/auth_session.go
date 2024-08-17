package auth

import (
	"context"
	"crypto/sha256"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type AuthSessions = map[string]*Session

func (as *SessionStore) NewAuthSession() (*Session, string) {
	expiresAt := time.Now().Add(168 * 4 * time.Hour)
	newToken := uuid.NewString()

	hash := sha256.New()
	hash.Write([]byte(newToken))
	ea := &Session{
		Token:        newToken,
		Expiry:       expiresAt,
		ActivateCode: as.Password,
		Type:         typeAuth,
	}
	return ea, newToken
}

func (as *SessionStore) GetAuthSession(token string) (*Session, bool) {
	as.sMu.Lock()
	defer as.sMu.Unlock()
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

func (as *SessionStore) AuthorizeAuth(next http.HandlerFunc) http.Handler {
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
		if !exists {
			log.Println("<Auth> Session doesn't exist")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}
		if s == nil {
			log.Println("<Auth> Session is nil")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		if r.Method == "POST" {
			log.Println(s.PostSuspendExpiry.Before(time.Now()))
			if !s.PostSuspendExpiry.IsZero() && !s.PostSuspendExpiry.Before(time.Now()) {
				log.Println("<Auth> User trying to use method POST, but is suspended")
				http.Error(w, "", http.StatusUnauthorized)
				return
			}
			s.PostSuspendExpiry = time.Now().Add(30 * time.Second)
		}
		if s.Type != typeAuth {
			log.Println("<Auth> Trying to log in AdminSession as AuthSession")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		if s.Expiry.Before(time.Now()) {
			log.Println("delete")
			delete(as.AuthSessions, sessionToken)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), &Session{}, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
