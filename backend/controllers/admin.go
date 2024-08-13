package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/views"
)

func (c controllers) panelController(w http.ResponseWriter, r *http.Request) {
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
		views.PanelPage(products, prices, c.sessionStore.Password).Render(r.Context(), w)
	}

}
func (c controllers) loginAdminController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		views.LoginPage().Render(r.Context(), w)
		return
	}
}
func (c controllers) productsController(w http.ResponseWriter, r *http.Request) {

	log.Println("POST PRODUCT")
	name := r.FormValue("name")
	description := r.FormValue("description")
	file, header, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	path, err := c.storage.AddImage(file, header)
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
	_, err = product.Add(c.data.Db, name, description, path)
	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}
	log.Println("Added a product")
	http.Redirect(w, r, "/panel", http.StatusSeeOther)

}
func (c controllers) productController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("POST PRODUCT")
		name := r.FormValue("name")
		description := r.FormValue("description")
		file, header, err := r.FormFile("image")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		path, err := c.storage.AddImage(file, header)
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
		_, err = product.Add(c.data.Db, name, description, path)
		if err != nil {
			http.Error(w, "Failed to add product", http.StatusInternalServerError)
			return
		}
		log.Println("Added a product")
		http.Redirect(w, r, "/panel", http.StatusSeeOther)
	}
	if r.Method == "PUT" {
		productId := r.PathValue("id")
		if productId == "" {
			http.Error(w, "not provided ID", http.StatusBadRequest)
			return
		}
		pId, err := strconv.ParseInt(productId, 10, 64)
		if err != nil {
			http.Error(w, "not provided ID", http.StatusBadRequest)
			return
		}
		p, err := c.data.Product.GetById(c.data.Db, int(pId))
		if err != nil {
			http.Error(w, "Failed to update a product", http.StatusInternalServerError)
			return
		}
		name := r.FormValue("name")
		if name != "" && name != p.Name {
			p.Name = name
		}
		description := r.FormValue("description")
		if description != "" && description != p.Description {
			p.Description = description
		}
		file, header, err := r.FormFile("image")
		if err != nil {
			if file != nil || header != nil {
				http.Error(w, "Failed to update a product, file err", http.StatusInternalServerError)
				return
			}
		}
		if file != nil {
			path, err := c.storage.AddImage(file, header)
			if err != nil {
				log.Println(err, "update")
				http.Error(w, "Failed to update a product, file upload err", http.StatusInternalServerError)
				return
			}
			p.ImageURL = path
		}
		err = c.data.Product.Update(c.data.Db, p)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to update a product", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (c controllers) pricesController(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("_method") == "PUT" {
		priceId := r.PathValue("id")
		if priceId == "" {
			http.Error(w, "not provided ID", http.StatusBadRequest)
			return
		}
		pId, err := strconv.ParseInt(priceId, 10, 64)
		if err != nil {
			http.Error(w, "not provided ID", http.StatusBadRequest)
			return
		}
		p, err := c.data.Price.GetById(c.data.Db, int(pId))
		if err != nil {
			log.Println(err)
			return
		}
		fp := r.FormValue("price")
		if fp == "" {
			http.Error(w, "not provided price", http.StatusBadRequest)
			return
		}
		priceFloat, err := strconv.ParseFloat(fp, 64)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		p.Price = priceFloat
		err = c.data.Price.UpdatePrice(c.data.Db, p)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/panel", http.StatusSeeOther)
	}
}
