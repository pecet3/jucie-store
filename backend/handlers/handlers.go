package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pecet3/my-api/auth"
	"github.com/pecet3/my-api/data"
)

type handlers struct {
	data data.Data
	ss   *auth.SessionStore
}

func Run(mux *http.ServeMux, d data.Data, ss *auth.SessionStore) {
	h := handlers{
		data: d,
		ss:   ss,
	}

	mux.HandleFunc("GET /api/products", h.handleProducts)
	mux.HandleFunc("GET /api/prices", h.handlePrices)
	mux.HandleFunc("POST /api/orders", h.handleOrders)
}

func (h handlers) handleProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.data.Product.GetAll(h.data.Db)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (h handlers) handlePrices(w http.ResponseWriter, r *http.Request) {
	prices, err := h.data.Price.GetAll(h.data.Db)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(prices)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (h handlers) handleOrders(w http.ResponseWriter, r *http.Request) {
	var order data.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	log.Println(order)
	w.WriteHeader(http.StatusOK)
}
