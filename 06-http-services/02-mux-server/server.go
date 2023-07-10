package main

import (
	"fmt"
	"net/http"

	"mux-server/controllers"

	"github.com/gorilla/mux"
)

// application specific handlers

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers details will be served")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {

	srv := mux.NewRouter()

	srv.HandleFunc("/", indexHandler)
	srv.HandleFunc("/customers", customersHandler)

	srv.HandleFunc("/products", controllers.GetAllProductsHandler).Methods("GET")
	srv.HandleFunc("/products/{id}", controllers.GetOneProductHandler).Methods("GET")
	srv.HandleFunc("/products", controllers.NewProductHandler).Methods("POST")

	// implement delete & put operations

	http.ListenAndServe(":8080", srv)
}
