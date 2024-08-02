package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/views"
	"github.com/pecet3/my-api/views/components"
)

func (c handlers) panelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		products, err := c.data.Product.GetAll(c.data.Db)
		if err != nil {
			http.Error(w, "products", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		prices, err := c.data.Price.GetAll(c.data.Db)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		views.PanelPage(products, prices).Render(r.Context(), w)
	}

}
func (c handlers) loginAdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.LoginPage().Render(r.Context(), w)
		return
	}
	if r.Method == "POST" {
		name := os.Getenv("USER_NAME")
		password := os.Getenv("USER_PASSWORD")

		formUser := r.FormValue("username")
		formPassword := r.FormValue("password")

		log.Println("New panel login")

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

func (c handlers) productsAdminHandler(w http.ResponseWriter, r *http.Request) {
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
		path, err := c.storage.AddImage(file, header, "/images")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error Saving or compressing a file", http.StatusInternalServerError)
			return
		}
		product := data.Product{
			Name:        name,
			Description: description,
			ImageURL:    path,
		}
		log.Println(path)
		_, err = product.Add(c.data.Db, name, description, path)
		if err != nil {
			http.Error(w, "Failed to add product", http.StatusInternalServerError)
			return
		}
		log.Println("Added a product")
		http.Redirect(w, r, "/panel", http.StatusSeeOther)
	}
}
