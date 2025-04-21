package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Order struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Status    string `json:"status"`
}

var orders = []Order{
	{ID: "1", ProductID: "1", Status: "completed"},
	{ID: "2", ProductID: "2", Status: "processing"},
}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/orders", GetOrders).Methods("GET")
	r.HandleFunc("/orders", CreateOrder).Methods("POST")

	return r
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	orders = append(orders, order)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
