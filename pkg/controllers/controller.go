package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/burakkarasel/grocery-shop-CRUD-API/pkg/models"
	"github.com/burakkarasel/grocery-shop-CRUD-API/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewProduct models.Product

// GetProducts handler gets all products from DB by using models package then we marshal and send as response
func GetProducts(w http.ResponseWriter, r *http.Request) {
	Products := models.GetAllProducts()
	res, _ := json.Marshal(Products)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetProductById handler takes the ID from request then find relevant product and send it after marshaling
func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing product ID: ", err)
	}

	productInfo, _ := models.GetProductById(ID)

	if productInfo.Name == "" && productInfo.Brand == "" && productInfo.Quantity == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		res, _ := json.Marshal(productInfo)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// CreateNewProduct handler parses the request body then creates a new product using models package then we marshal
//and return it
func CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	CreateProduct := &models.Product{}
	utils.ParseRequestBody(r, CreateProduct)
	b := CreateProduct.CreateNewProduct()

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// DeleteProductById handler gets product's ID and deletes it from DB then marshal the product and send as response
func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing product ID: ", err)
	}

	product := models.DeleteProduct(ID)
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateProductById handler parses request body then finds the old record checks if all new info is not empty and then
// we update DB according to the new informations and send same new informations after marshaling as response
func UpdateProductById(w http.ResponseWriter, r *http.Request) {
	var updateProduct = &models.Product{}
	utils.ParseRequestBody(r, updateProduct)

	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing product ID: ", err)
	}

	productInfo, db := models.GetProductById(ID)

	if updateProduct.Name != "" {
		productInfo.Name = updateProduct.Name
	}
	if updateProduct.Brand != "" {
		productInfo.Brand = updateProduct.Brand
	}
	if updateProduct.Quantity != "" {
		productInfo.Quantity = updateProduct.Quantity
	}

	// Saved updated product info
	db.Save(&productInfo)

	res, _ := json.Marshal(productInfo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
