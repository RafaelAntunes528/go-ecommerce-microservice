package app

import (
	"github.com/RafaelAntunes528/go-ecommerce/product-service/internal/handlers"
	"github.com/gorilla/mux"
)

type ProductApp struct {
	Router *mux.Router
}

func NewProductApp() *ProductApp {
	app := &ProductApp{
		Router: mux.NewRouter(),
	}
	app.setupRoutes()
	return app
}

func (a *ProductApp) setupRoutes() {
	a.Router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	a.Router.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
}
