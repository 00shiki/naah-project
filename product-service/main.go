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
	// Admin routes
	admin := e.Group("/admin")
	admin.POST("/shoe-details", handlers.CreateShoeDetail)       // Create shoe detail
	admin.GET("/shoe-details", handlers.GetShoeDetails)          // Get all shoe details
	admin.GET("/shoe-details/:id", handlers.GetShoeDetailByID)   // Get shoe detail by ID
	admin.PUT("/shoe-details/:id", handlers.UpdateShoeDetail)    // Update shoe detail
	admin.DELETE("/shoe-details/:id", handlers.DeleteShoeDetail) // Delete shoe detail

	// Customer routes
	customer := e.Group("/customer")
	customer.GET("/shoe-details", handlers.GetShoeDetailsForCustomer)        // Get shoe details for customer
	customer.GET("/shoe-details/:id", handlers.GetShoeDetailForCustomerByID) // Get specific shoe detail for customer

	// Start the server
	e.Logger.Fatal(e.Start(":50052"))
}
