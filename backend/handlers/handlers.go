package handlers

import (
	"net/http"

	"github.com/pecet3/my-api/auth"
	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/storage"
)

type handlers struct {
	data         data.Data
	storage      storage.StorageServices
	sessionStore *auth.SessionStore
}

func Run(mux *http.ServeMux, d data.Data, s storage.StorageServices, ss *auth.SessionStore) {
	c := handlers{
		data:         d,
		storage:      s,
		sessionStore: ss,
	}

	// mux.Handle("/panel", ss.AuthorizeAdmin(c.panelHandler))

	// mux.Handle("/products/{id}", ss.AuthorizeAdmin(c.productsAdminHandler))
	// mux.Handle("/prices/{id}", ss.AuthorizeAdmin(c.pricesHandler))
	mux.HandleFunc("/panel", c.panelHandler)
	mux.HandleFunc("/products/{id}", c.productsAdminHandler)
	mux.HandleFunc("/prices/{id}", c.pricesHandler)
	mux.HandleFunc("/login", c.loginAdminHandler)
}
