package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

var products = []Product{
	{Id: 101, Name: "Pen", Cost: 5},
	{Id: 102, Name: "Pencil", Cost: 2},
}

type AppServer struct {
}

func (as *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			newProduct.Id = 100 + len(products) + 1
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "All the customers will be served")
	default:
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}

}

func main() {
	server := &AppServer{}
	http.ListenAndServe(":8080", server)
}

/*
GET - /products
GET - /customers
*/
