package controllers

import (
	"net/http"

	"github.com/pecet3/my-api/auth"
	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/storage"
	"github.com/pecet3/my-api/views_page"
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
	mux.HandleFunc("/", c.mainPageController)
	mux.HandleFunc("/how", c.howPageController)

	mux.Handle("/panel", ss.Authorize(c.panelController))
	mux.HandleFunc("GET /products", c.productsController)
	mux.Handle("POST /products", ss.Authorize(c.productsController))

	mux.HandleFunc("/login", c.loginController)
}

func (c controllers) mainPageController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views_page.MainPage().Render(r.Context(), w)
	}
}

func (c controllers) howPageController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views_page.HowPage().Render(r.Context(), w)
	}
}
