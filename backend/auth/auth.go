package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/utils"
)

type auth struct {
	ss   *SessionStore
	data data.Data
}

type LoginDto struct {
	Password string `json:"password"`
}

func Run(srv *http.ServeMux, ss *SessionStore, data data.Data) {
	a := &auth{
		ss:   ss,
		data: data,
	}

	srv.HandleFunc("/auth/login-admin", a.handleAdminLogin)
	srv.HandleFunc("/auth/login", a.handleLogin)

	go changePasswordLoop(ss)
	go cleanUpExpiredSessionsLoop(ss)
}
func (a auth) handleLogin(w http.ResponseWriter, r *http.Request) {
	currentPswd := a.ss.GetCurrentPassword()
	formPswd := r.FormValue("password")
	log.Printf("<Auth> User with IP:%s tries to login, currentPswd: %s, formPswd: %s",
		utils.GetIP(r), currentPswd, formPswd)

	if currentPswd == formPswd {
		us, token := a.ss.NewAuthSession()
		a.ss.AddAuthSession(token, us)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Expires:  us.Expiry,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (a auth) handleAdminLogin(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("USER_NAME")
	password := os.Getenv("USER_PASSWORD")

	formUser := r.FormValue("username")
	formPassword := r.FormValue("password")

	if name == formUser && password == formPassword {
		us, token := a.ss.NewAdminSession(r)
		a.ss.AddAdminSession(token, us)
		http.SetCookie(w, &http.Cookie{
			Name:     "admin_token",
			Value:    token,
			Expires:  us.Expiry,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})
		http.Redirect(w, r, "/panel", http.StatusSeeOther)
		return
	}
	http.Error(w, "wrong credentials", http.StatusUnauthorized)
}
