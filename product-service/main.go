package main

import (
	"log"
	"net/http"

	"github.com/RafaelAntunes528/go-ecommerce/product-service/internal/app"
)

func main() {
	app := app.NewProductApp()
	log.Println("Product service starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", app.Router))
}
