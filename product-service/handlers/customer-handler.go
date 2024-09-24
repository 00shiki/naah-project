package handlers

import (
	"log"
	"net/http"
	"product-service/config"
	"product-service/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// show
func GetProductsForCustomer(c echo.Context) error {
	log.Println("Fetching all products")
	var products []models.ShoeModel
	if err := config.DB.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch products", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

// show by id
func GetProductForCustomerByID(c echo.Context) error {
	log.Println("Fetching product by ID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid product ID"})
	}

	var product models.ShoeModel
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}
