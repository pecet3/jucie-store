package auth

import (
	"log"
	"net/http"
	"time"
)

func (a *auth) handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	existingSession, exists := a.sessionStore.GetAuthSession(sessionToken)
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if existingSession.Token != sessionToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !existingSession.Expiry.Before(time.Now()) {
		log.Println("User wants to reissue a new token, but has token still active")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	a.sessionStore.RemoveAuthSession(sessionToken)

	// todo: check if email exists

	us, token := a.sessionStore.NewAuthSession(r, existingSession.UserId)
	a.sessionStore.AddAuthSession(token, us)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: us.Expiry,
	})
	w.WriteHeader(http.StatusAccepted)
}
