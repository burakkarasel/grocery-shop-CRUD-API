package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// ConnectToDB is connecting our DB using GORM
func ConnectToDB() {
	d, err := gorm.Open("mysql", "*@/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB returns the DB we connected earlier
func GetDB() *gorm.DB {
	return db
}
