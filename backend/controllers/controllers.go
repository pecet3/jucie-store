package controllers

import (
	"net/http"

	"github.com/pecet3/my-api/auth"
	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/storage"
	"github.com/pecet3/my-api/views"
)

type controllers struct {
	data         data.Data
	storage      storage.StorageServices
	sessionStore *auth.SessionStore
}

func Run(mux *http.ServeMux, d data.Data, s storage.StorageServices, ss *auth.SessionStore) {
	c := controllers{
		data:         d,
		storage:      s,
		sessionStore: ss,
	}

	mux.Handle("/panel", ss.AuthorizeAdmin(c.panelController))
	mux.Handle("/products/{id}", ss.AuthorizeAdmin(c.productsAdminController))
	mux.Handle("/prices/{id}", ss.AuthorizeAdmin(c.pricesController))

	// mux.HandleFunc("/panel", c.panelController)
	// mux.HandleFunc("/products/{id}", c.productsAdminController)
	// mux.HandleFunc("/prices/{id}", c.pricesController)

	mux.HandleFunc("/login-admin", c.loginAdminController)
	mux.HandleFunc("/login", c.userLoginController)
}

func (c controllers) userLoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.EntryPage().Render(r.Context(), w)
	}
}
