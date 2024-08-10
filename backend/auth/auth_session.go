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
	expiresAt := time.Now().Add(1 * time.Minute)
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
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}
		if s.Type != typeAuth {
			log.Println("<Auth> Trying to log in AdminSession as AuthSession")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), &Session{}, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
