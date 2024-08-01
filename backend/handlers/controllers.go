package handlers

import (
	"net/http"

	"github.com/pecet3/my-api/auth"
	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/storage"
	"github.com/pecet3/my-api/views_page"
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
	mux.HandleFunc("/", c.mainPageHandler)
	mux.HandleFunc("/how", c.howPageHandler)

	mux.Handle("/panel", ss.AuthorizeAdmin(c.panelHandler))
	mux.HandleFunc("GET /products", c.productsHandler)
	mux.Handle("POST /products", ss.AuthorizeAdmin(c.productsHandler))

	mux.HandleFunc("/login", c.loginAdminHandler)
}

func (c handlers) mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views_page.MainPage().Render(r.Context(), w)
	}
}

func (c handlers) howPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views_page.HowPage().Render(r.Context(), w)
	}
}
