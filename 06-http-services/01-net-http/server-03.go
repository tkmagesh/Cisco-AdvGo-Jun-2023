package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

/*
	type AppServer struct {
		routes map[string]http.HandlerFunc
	}

	func NewAppServer() *AppServer {
		return &AppServer{
			routes: make(map[string]http.HandlerFunc),
		}
	}

	func (as *AppServer) Register(url string, handler http.HandlerFunc) {
		as.routes[url] = handler
	}

	func (as *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
		handler, ok := as.routes[r.URL.Path]
		if !ok {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		handler(w, r)
	}
*/

// middlewares
func logHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s -%s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func profileHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

// request handlers
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

func main() {
	server := http.NewServeMux()
	/*
		server.HandleFunc("/", IndexHandler)
		server.HandleFunc("/products", ProductsHandler)
		server.HandleFunc("/customers", CustomersHandler)
	*/
	server.HandleFunc("/", profileHandler(logHandler(IndexHandler)))
	server.HandleFunc("/products", profileHandler(logHandler(ProductsHandler)))
	server.HandleFunc("/customers", profileHandler(logHandler(CustomersHandler)))

	http.ListenAndServe(":8080", server)
}

/*
GET - /products
GET - /customers
*/
