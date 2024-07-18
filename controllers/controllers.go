package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/storage"
	"github.com/pecet3/my-api/views"
	"github.com/pecet3/my-api/views/components"
)

type controllers struct {
	data    data.Data
	storage storage.StorageServices
}

func Run(mux *http.ServeMux, d data.Data, s storage.StorageServices) {
	c := controllers{
		data:    d,
		storage: s,
	}
	mux.HandleFunc("/panel", c.panelController)
	mux.HandleFunc("/products", c.productsController)
	mux.HandleFunc("/login", c.loginController)
	mux.HandleFunc("/register", c.registerController)

}

func (c controllers) panelController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.PanelPage().Render(r.Context(), w)

	}

}
func (c controllers) loginController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.LoginPage().Render(r.Context(), w)

	}

}
func (c controllers) registerController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.PanelPage().Render(r.Context(), w)

	}

}
func (c controllers) productsController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		products, err := c.data.Product.GetAll(c.data.Db)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		components.ProductsDisplay(products).Render(r.Context(), w)

	}
	if r.Method == "POST" {
		// Parse form values
		name := r.FormValue("name")
		description := r.FormValue("description")
		file, header, err := r.FormFile("image")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}

		path, err := c.storage.AddImage(file, header, "")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error Saving or compressing a file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File uploaded successfully: %s", path)
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		category := r.FormValue("category")

		product := data.Product{
			Name:        name,
			Description: description,
			ImageURL:    path,
			Quantity:    quantity,
			Price:       price,
			Category:    category,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		_, err = product.Add(c.data.Db, name, description, path, quantity, price, category)
		if err != nil {
			http.Error(w, "Failed to add product", http.StatusInternalServerError)
			return
		}

		// Redirect to the success page or product list
		http.Redirect(w, r, "/panel", http.StatusSeeOther)
	}
}
