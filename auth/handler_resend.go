package auth

import (
	"net/http"
)

func (a *auth) handleResendActivateEmail(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No cookie!!", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	acode := r.PathValue("id")
	token := cookie.Value
	ea, exists := a.sessionStore.GetEmailSession(token)
	if !exists {
		http.Error(w, "Unauthorized 1", http.StatusUnauthorized)
		return
	}
	if ea.Token != token || acode != ea.ActivateCode {
		http.Error(w, "Unauthorized 2", http.StatusUnauthorized)
		return
	}
	user, err := a.userServices.GetById(a.db, ea.UserId)
	if user == nil || err != nil {
		http.Error(w, "Unauthorized 3", http.StatusUnauthorized)
		return
	}
	if user.IsActive {
		http.Error(w, "Unauthorized 4", http.StatusUnauthorized)
		return
	}
	a.sessionStore.RemoveEmailSession(token)
	newSession, newToken := a.sessionStore.NewEmailSession(r, user.Email, ea.UserId)
	a.sessionStore.AddEmailSession(newToken, newSession)
	if err = a.emailServices.SendActivateEmail(user.Email, newSession.ActivateCode, user.Name); err != nil {
		http.Error(w, "Unauthorized 5", http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: newSession.Expiry,
	})
	w.WriteHeader(http.StatusAccepted)
}
