package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"price"`
}

var products = []Product{
	{Id: 101, Name: "Pen", Cost: 10},
	{Id: 102, Name: "Pencil", Cost: 5},
	{Id: 103, Name: "Marker", Cost: 50},
}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
func GetOneProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	for _, p := range products {
		if p.Id == requestedId {
			if err := json.NewEncoder(w).Encode(p); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "product not found", http.StatusNotFound)
}

func NewProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
	}
	newProduct.Id = len(products) + 101
	products = append(products, newProduct)
	if err := json.NewEncoder(w).Encode(newProduct); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
