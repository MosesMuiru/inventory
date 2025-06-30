package main

import (
	db "github.com/inventory/storage"
	"github.com/inventory/storage/models"
)

func main() {

	db.ConnectDb()

	DB := db.DB

	// autho migrate models

	DB.AutoMigrate(&models.Store{})
	DB.AutoMigrate(&models.User{})

}
