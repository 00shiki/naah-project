package handlers

import (
	"database/sql"
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
	query := "SELECT model_id, name, price FROM shoe_models"
	rows, err := config.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch products", "error": err.Error()})
	}
	defer rows.Close()

	var products []models.ShoeModel
	for rows.Next() {
		var product models.ShoeModel
		if err := rows.Scan(&product.ModelID, &product.Name, &product.Price); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not scan product", "error": err.Error()})
		}
		query2 := "SELECT shoe_id, size, stock FROM shoe_details WHERE model_id = ?"
		rows2, err := config.DB.Query(query2, product.ModelID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch product", "error": err.Error()})
		}
		shoeDetails := make([]models.ShoeDetail, 0)
		for rows2.Next() {
			var shoeDetail models.ShoeDetail
			if err := rows2.Scan(&shoeDetail.ShoeID, &shoeDetail.Size, &shoeDetail.Stock); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch product", "error": err.Error()})
			}
			shoeDetails = append(shoeDetails, shoeDetail)
		}
		product.ShoeDetails = shoeDetails
		products = append(products, product)
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

	query := "SELECT model_id, name, price FROM shoe_models WHERE model_id = ?"
	row := config.DB.QueryRow(query, id)

	var product models.ShoeModel
	if err := row.Scan(&product.ModelID, &product.Name, &product.Price); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch product", "error": err.Error()})
	}
	query = "SELECT shoe_id, size, stock FROM shoe_details WHERE model_id = ?"
	rows, err := config.DB.Query(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch product", "error": err.Error()})
	}
	shoeDetails := make([]models.ShoeDetail, 0)
	for rows.Next() {
		var shoeDetail models.ShoeDetail
		if err := rows.Scan(&shoeDetail.ShoeID, &shoeDetail.Size, &shoeDetail.Stock); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch product", "error": err.Error()})
		}
		shoeDetails = append(shoeDetails, shoeDetail)
	}
	product.ShoeDetails = shoeDetails
	return c.JSON(http.StatusOK, product)
}
