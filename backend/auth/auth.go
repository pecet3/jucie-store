package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/pecet3/my-api/data"
)

type auth struct {
	ss   *SessionStore
	data data.Data
}

func Run(srv *http.ServeMux, ss *SessionStore, data data.Data) {
	a := &auth{
		ss:   ss,
		data: data,
	}

	srv.HandleFunc("/auth/admin-login", a.AdminLogin)
}

func (a auth) AdminLogin(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("USER_NAME")
	password := os.Getenv("USER_PASSWORD")

	formUser := r.FormValue("username")
	formPassword := r.FormValue("password")

	log.Println("New panel login")

	if name == formUser && password == formPassword {
		us, token := a.ss.NewAdminSession(r, 123)
		a.ss.AddAdminSession(token, us)
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   token,
			Expires: us.Expiry,
		})
		http.Redirect(w, r, "/panel", http.StatusSeeOther)
		return
	}
	http.Error(w, "wrong credentials", http.StatusUnauthorized)
}
