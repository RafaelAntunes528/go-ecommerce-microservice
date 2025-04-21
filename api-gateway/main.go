package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	productURL, _ := url.Parse("http://localhost:8081")
	orderURL, _ := url.Parse("http://localhost:8082")

	r.PathPrefix("/products").Handler(httputil.NewSingleHostReverseProxy(productURL))
	r.PathPrefix("/orders").Handler(httputil.NewSingleHostReverseProxy(orderURL))

	log.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
