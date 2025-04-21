package main

import (
	"log"
	"net/http"

	"github.com/RafaelAntunes528/go-ecommerce/order-service/internal/app"
)

func main() {
	app := app.NewOrderApp()
	log.Println("Order service starting on :8082")
	log.Fatal(http.ListenAndServe(":8082", app.Router))
}
