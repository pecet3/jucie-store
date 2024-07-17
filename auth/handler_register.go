package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func (a *auth) handleRegister(w http.ResponseWriter, r *http.Request) {
	var regUser registerDto
	if err := json.NewDecoder(r.Body).Decode(&regUser); err != nil {
		http.Error(w, "Bad json request", http.StatusBadRequest)
		return
	}
	err := a.validator.Struct(regUser)
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
	isNameUnsafe, err := regexp.MatchString(`[!@#~$%^&*(),.?":{}|<>]`, regUser.Name)
	if err != nil {
		log.Printf("Unexpected error during validation name: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isNameUnsafe {
		log.Printf("User provided unsafe name during register process: %v", regUser)
		http.Error(w, "Name cannot contains the special characters", http.StatusBadRequest)
		return
	}
	isValidPswd, err := a.pswdServices.validatePassword(regUser.Password)
	if err != nil || !isValidPswd {
		log.Printf("User provided invalid password during register process: %v", regUser.Email)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	}

	existingUser, err := a.userServices.GetByEmail(a.db, regUser.Email)
	if err == nil || existingUser != nil {
		log.Println("User wanted to create an account with existing email: ", regUser.Email)
		http.Error(w, "User with provided email exists", http.StatusBadRequest)
		return
	}

	salt, err := a.pswdServices.generateSalt()
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	hashedPswd, err := a.pswdServices.hashPassword(salt, regUser.Password)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	userId, err := a.userServices.Add(a.db, regUser.Name, regUser.Email, hashedPswd, salt)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	ea, token := a.sessionStore.NewEmailSession(r, regUser.Email, int(userId))
	a.sessionStore.AddEmailSession(token, ea)
	url := "localhost:8090/auth/activate-account/" + ea.ActivateCode

	log.Printf("Created an user with id:%v name: %s and email: %s", userId, regUser.Name, regUser.Email)
	err = a.emailServices.SendActivateEmail(regUser.Email, url, regUser.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	log.Printf("Email to: %s with activate link sent", regUser.Email)

	http.SetCookie(w, &http.Cookie{
		Name:    "email_session",
		Value:   token,
		Expires: ea.Expiry,
	})
	log.Printf("Register Success, email:%v", *ea)
	w.WriteHeader(http.StatusAccepted)
}
