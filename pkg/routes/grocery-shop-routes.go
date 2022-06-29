package routes

import (
	"github.com/burakkarasel/grocery-shop-CRUD-API/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterGroceryShopRoutes = func(router *mux.Router) {
	router.HandleFunc("/products/", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/product/", controllers.CreateNewProduct).Methods("POST")
	router.HandleFunc("/product/{productId}", controllers.DeleteProductById).Methods("DELETE")
	router.HandleFunc("/product/{productId}", controllers.UpdateProductById).Methods("PUT")
}
