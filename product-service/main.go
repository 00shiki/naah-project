package main

import (
	"product-service/config"
	"product-service/handlers"

	"github.com/labstack/echo/v4"
)

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
