package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/misael28/GoApi/configs"
	"github.com/misael28/GoApi/internal/entity"
	"github.com/misael28/GoApi/internal/infra/database"
	"github.com/misael28/GoApi/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	r:= chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProdcut)

	http.ListenAndServe(":8000", r)
}