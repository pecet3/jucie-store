package auth

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pecet3/my-api/data"
)

type auth struct {
	db            *sql.DB
	pswdServices  passwordServices
	sessionStore  *SessionStore
	emailServices emailServices
	validator     *validator.Validate
	userServices  data.User
}

func Run(srv *http.ServeMux, ss *SessionStore, d data.Data, v *validator.Validate) {

	a := &auth{
		db:            d.Db,
		pswdServices:  &password{},
		sessionStore:  ss,
		validator:     v,
		emailServices: &email{},
		userServices:  d.User,
	}

	u, _ := a.userServices.GetById(d.Db, 1)
	log.Println(u)

	srv.HandleFunc("POST /auth/register", a.handleRegister)
	srv.HandleFunc("POST /auth/login", a.handleLogin)
	srv.HandleFunc("POST /auth/refresh-token", a.handleRefreshToken)
	srv.HandleFunc("GET /auth/activate-account/{id}", a.handleActivateAccount)
	srv.HandleFunc("GET /auth/resend-email", a.handleResendActivateEmail)
	srv.Handle("GET /hello", ss.Authorize(handleHello))

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value("session_token")
	idStr := r.PathValue("id")
	w.Write([]byte(idStr))
	fmt.Fprintf(w, "Hello %s", session)
	w.Write([]byte(idStr))
}
