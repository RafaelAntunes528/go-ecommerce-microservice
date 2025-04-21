package app

import (
	"github.com/RafaelAntunes528/go-ecommerce/order-service/internal/handlers"
	"github.com/gorilla/mux"
)

type OrderApp struct {
	Router *mux.Router
}

func NewOrderApp() *OrderApp {
	app := &OrderApp{
		Router: mux.NewRouter(),
	}
	app.setupRoutes()
	return app
}

func (a *OrderApp) setupRoutes() {
	a.Router.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	a.Router.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
}
