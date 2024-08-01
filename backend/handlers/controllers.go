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

	mux.Handle("/panel", ss.AuthorizeAdmin(c.panelHandler))
	mux.Handle("/products", ss.AuthorizeAdmin(c.productsAdminHandler))

	mux.HandleFunc("/login", c.loginAdminHandler)
}
