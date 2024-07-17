package controllers

import (
	"net/http"

	"github.com/pecet3/my-api/views"
)

func Run(mux *http.ServeMux) {
	mux.HandleFunc("/test", testController)
}

func testController(w http.ResponseWriter, r *http.Request) {
	views.Hello().Render(r.Context(), w)
	views.Hello().Render(r.Context(), w)

}
