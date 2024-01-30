package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/osmar-vil/goapi/internal/database"
	"github.com/osmar-vil/goapi/internal/service"
	"github.com/osmar-vil/goapi/internal/webserver"
)

func main() {
	db, e := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if e != nil {
		panic(e.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()

	// MIDDLEWARE
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	// CATEGORIES
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	// PRODUCTS
	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/category/{id}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("SERVER RUN ON :8080")
	http.ListenAndServe(":8080", c)
}
