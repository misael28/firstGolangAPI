package main

import (
	"net/http"

	"github.com/misael28/GoApi/configs"
	"github.com/misael28/GoApi/internal/entity"
	"github.com/misael28/GoApi/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/misael28/GoApi/internal/infra/webserver/handlers"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProdcut)
	http.ListenAndServe(":8000", nil)
}