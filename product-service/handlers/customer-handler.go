package handlers

import (
	"log"
	"net/http"
	"product-service/config"
	"product-service/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetProductsForCustomer godoc
// @Summary Fetch all products for customers
// @Description Retrieve a list of all available products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.ShoeModel
// @Failure 500 {object} map[string]string
// @Router /products [get]
func GetProductsForCustomer(c echo.Context) error {
	log.Println("Fetching all products")
	var products []models.ShoeModel
	if err := config.DB.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch products", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

// GetProductForCustomerByID godoc
// @Summary Fetch a product by ID for customers
// @Description Retrieve a specific product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.ShoeModel
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
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
