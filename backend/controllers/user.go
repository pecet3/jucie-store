package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/pecet3/my-api/utils"
	"github.com/pecet3/my-api/views"
)

func (c controllers) userLoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.EntryPage().Render(r.Context(), w)
	}
}

func (c controllers) serveReact(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./static/dist"))
	path := r.URL.Path
	_, err := os.Stat("./static/dist" + path)

	if os.IsNotExist(err) {
		http.ServeFile(w, r, "./static/dist/index.html")
		return
	}
	log.Printf("<Controllers> User with IP:%s entered the protected react app", utils.GetIP(r))

	fs.ServeHTTP(w, r)
}
