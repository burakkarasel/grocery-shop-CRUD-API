package main

import (
	"github.com/burakkarasel/grocery-shop-CRUD-API/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterGroceryShopRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
