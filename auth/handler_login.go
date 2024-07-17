package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (a *auth) handleLogin(w http.ResponseWriter, r *http.Request) {
	var creds credentialsDto
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = a.validator.Struct(creds)
	if err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Printf("Unexpected error during validation: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.Error(w, fmt.Sprintf("%v", errors), http.StatusBadRequest)
		return
	}

	u, err := a.userServices.GetByEmail(a.db, creds.Email)
	log.Println(u)
	if u.Email != creds.Email || err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	isValidPswd, err := a.pswdServices.verifyPassword(u.Password, u.Salt)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	if !isValidPswd {
		log.Printf("User with email:%s provided a wrong password", u.Email)
		http.Error(w, "invalid email or/and password", http.StatusBadRequest)
		return
	}
	if !u.IsActive {
		log.Printf("User with email:%s provided a wrong password", u.Email)
		http.Error(w, "Confirm your account, check your email", http.StatusBadRequest)
		return
	}

	us, token := a.sessionStore.NewAuthSession(r, u.Id)
	a.sessionStore.AddAuthSession(token, us)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: us.Expiry,
	})
	log.Printf("Login Success, session:%v", *us)
	w.WriteHeader(http.StatusAccepted)
}
