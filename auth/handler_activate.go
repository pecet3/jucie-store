package auth

import (
	"log"
	"net/http"
)

func (a *auth) handleActivateAccount(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email_session")
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
	log.Println(exists)
	if !exists {
		log.Println(exists)
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
	a.sessionStore.RemoveEmailSession(token)
	query := "update users set is_active = $2 where id = $1"

	_, err = a.db.Exec(query, ea.UserId, true)

	if err != nil {
		log.Println(err)
		http.Error(w, "sql error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
