package main

import (
	"product-service/config"
	"product-service/handlers"

	"github.com/labstack/echo/v4"
)

// @title CRUD NAAH PROJECT API
// @version 1.0.0
// @description This is the API documentation for the CRUD NAAH PROJECT, a comprehensive system designed to manage shoe models and customer interactions.
// The API provides endpoints for administrators to create, read, update, and delete shoe models, as well as for customers to browse available products.
// Features include user-friendly interfaces for easy management of shoe inventory and seamless access to product information for customers.
// @host localhost:50052
// @BasePath /api/v1

func main() {
	e := echo.New()

	// Initialize the database
	config.InitDB()

	// Admin routes
	admin := e.Group("/admin")
	admin.POST("/shoe-models", handlers.CreateShoeModel)
	admin.GET("/shoe-models", handlers.GetShoeModels)
	admin.GET("/shoe-models/:id", handlers.GetShoeModelByID)
	admin.PUT("/shoe-models/:id", handlers.UpdateShoeModel)
	admin.DELETE("/shoe-models/:id", handlers.DeleteShoeModel)

	// Customer routes
	customer := e.Group("/customer")
	customer.GET("/products", handlers.GetProductsForCustomer)
	customer.GET("/products/:id", handlers.GetProductForCustomerByID)

	// Start the server
	e.Logger.Fatal(e.Start(":50052"))
}
