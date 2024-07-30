package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/views"
	"github.com/pecet3/my-api/views/components"
)

func (c controllers) panelController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.PanelPage().Render(r.Context(), w)
	}

}
func (c controllers) loginController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.LoginPage().Render(r.Context(), w)
		return
	}
	if r.Method == "POST" {
		name := os.Getenv("USER_NAME")
		password := os.Getenv("USER_PASSWORD")

		formUser := r.FormValue("username")
		formPassword := r.FormValue("password")

		log.Println(formUser, formPassword)

		if name == formUser && password == formPassword {
			us, token := c.sessionStore.NewAuthSession(r, 123)
			c.sessionStore.AddAuthSession(token, us)
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   token,
				Expires: us.Expiry,
			})
			http.Redirect(w, r, "/panel", http.StatusSeeOther)
			return
		}
		http.Error(w, "wrong credentials", http.StatusUnauthorized)
		return
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
		log.Println("Added a product")
		http.Redirect(w, r, "/panel", http.StatusSeeOther)
	}
}
