package models

import (
	"github.com/burakkarasel/grocery-shop-CRUD-API/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Product is the struct for our products in the grocery shop
type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Brand    string `json:"brand"`
	Quantity string `json:"quantity"`
}

// init function connects and gets the DB using config package
func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

// GetAllProducts returns matching rows in our database
func GetAllProducts() []Product {
	var Products []Product
	db.Find(&Products)
	return Products
}

// GetProductById returns the product according to given ID and returns it
func GetProductById(Id int64) (*Product, *gorm.DB) {
	var searchedProduct Product
	db := db.Where("ID=?", Id).Find(&searchedProduct)
	return &searchedProduct, db
}

// CreateNewProduct is a receiver func for pointer of product struct it creates a new record in DB and returns product
func (p *Product) CreateNewProduct() *Product {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

// DeleteProduct deletes the Product's record in DB and returns the product
func DeleteProduct(Id int64) Product {
	var product Product
	db.Where("ID=?", Id).Delete(product)
	return product
}
