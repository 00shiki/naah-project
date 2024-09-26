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

// GetShoeDetailsForCustomer godoc
// @Summary Fetch all shoe details for customers
// @Description Retrieve a list of all available shoe details
// @Tags shoe-details
// @Accept json
// @Produce json
// @Success 200 {array} models.ShoeDetail
// @Failure 500 {object} map[string]string
// @Router /customer/shoe-details [get]
func GetShoeDetailsForCustomer(c echo.Context) error {
	log.Println("Fetching all shoe details")
	query := `
		SELECT d.shoe_id, d.model_id, d.size, d.stock, m.name, m.price 
		FROM shoe_details d 
		JOIN shoe_models m ON d.model_id = m.model_id`
	rows, err := config.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch shoe details", "error": err.Error()})
	}
	defer rows.Close()

	var shoeDetails []struct {
		models.ShoeDetail
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	for rows.Next() {
		var shoeDetail struct {
			models.ShoeDetail
			Name  string `json:"name"`
			Price int    `json:"price"`
		}
		if err := rows.Scan(&shoeDetail.ShoeID, &shoeDetail.ModelID, &shoeDetail.Size, &shoeDetail.Stock, &shoeDetail.Name, &shoeDetail.Price); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not scan shoe detail", "error": err.Error()})
		}
		shoeDetails = append(shoeDetails, shoeDetail)
	}
	return c.JSON(http.StatusOK, shoeDetails)
}

// GetShoeDetailForCustomerByID godoc
// @Summary Fetch a shoe detail by ID for customers
// @Description Retrieve a specific shoe detail by its ID
// @Tags shoe-details
// @Accept json
// @Produce json
// @Param id path int true "Shoe Detail ID"
// @Success 200 {object} models.ShoeDetail
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /customer/shoe-details/{id} [get]
func GetShoeDetailForCustomerByID(c echo.Context) error {
	log.Println("Fetching shoe detail by ID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid shoe detail ID"})
	}

	query := `
		SELECT d.shoe_id, d.model_id, d.size, d.stock, m.name, m.price 
		FROM shoe_details d 
		JOIN shoe_models m ON d.model_id = m.model_id 
		WHERE d.shoe_id = ?`
	row := config.DB.QueryRow(query, id)

	var shoeDetail struct {
		models.ShoeDetail
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	if err := row.Scan(&shoeDetail.ShoeID, &shoeDetail.ModelID, &shoeDetail.Size, &shoeDetail.Stock, &shoeDetail.Name, &shoeDetail.Price); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Shoe detail not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch shoe detail", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, shoeDetail)
}
