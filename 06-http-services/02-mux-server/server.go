package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

// middlewares
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func profileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Println(elapsed)
	})
}
func main() {

	srv := mux.NewRouter()
	srv.Use(logMiddleware)
	srv.Use(profileMiddleware)
	srv.HandleFunc("/", indexHandler)
	srv.HandleFunc("/customers", customersHandler)

	srv.HandleFunc("/products", controllers.GetAllProductsHandler).Methods("GET")
	srv.HandleFunc("/products/{id}", controllers.GetOneProductHandler).Methods("GET")
	srv.HandleFunc("/products", controllers.NewProductHandler).Methods("POST")

	// implement delete & put operations

	http.ListenAndServe(":8080", srv)
}
