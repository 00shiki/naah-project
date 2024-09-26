package handlers

import (
	"database/sql"
	"net/http"
	"product-service/config"
	"product-service/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetShoeDetails godoc
// @Summary      Get all shoe details
// @Description  Retrieve all shoe details including shoe_id, model_id, size, stock, name, and price
// @Tags         shoe-details
// @Accept       json
// @Produce      json
// @Success      200 {array} map[string]interface{}
// @Failure      500 {object} map[string]string
// @Router       /shoe-details [get]
func GetShoeDetails(c echo.Context) error {
	query := `
		SELECT sd.shoe_id, sd.model_id, sd.size, sd.stock, sm.name, sm.price 
		FROM shoe_details sd 
		JOIN shoe_models sm ON sd.model_id = sm.model_id
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch shoe details"})
	}
	defer rows.Close()

	var shoeDetails []map[string]interface{}
	for rows.Next() {
		var shoeDetail models.ShoeDetail
		var name string
		var price int
		if err := rows.Scan(&shoeDetail.ShoeID, &shoeDetail.ModelID, &shoeDetail.Size, &shoeDetail.Stock, &name, &price); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan shoe details"})
		}
		shoeDetails = append(shoeDetails, map[string]interface{}{
			"shoe_id":  shoeDetail.ShoeID,
			"model_id": shoeDetail.ModelID,
			"size":     shoeDetail.Size,
			"stock":    shoeDetail.Stock,
			"name":     name,
			"price":    price,
		})
	}

	return c.JSON(http.StatusOK, shoeDetails)
}

// GetShoeDetailByID godoc
// @Summary      Get a shoe detail by ID
// @Description  Retrieve a specific shoe detail by shoe_id including name and price
// @Tags         shoe-details
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Shoe ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /shoe-details/{id} [get]
func GetShoeDetailByID(c echo.Context) error {
	shoeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid shoe ID"})
	}

	var shoeDetail models.ShoeDetail
	var name string
	var price int
	query := `
		SELECT sd.shoe_id, sd.model_id, sd.size, sd.stock, sm.name, sm.price 
		FROM shoe_details sd 
		JOIN shoe_models sm ON sd.model_id = sm.model_id 
		WHERE sd.shoe_id = ?
	`
	err = config.DB.QueryRow(query, shoeID).Scan(&shoeDetail.ShoeID, &shoeDetail.ModelID, &shoeDetail.Size, &shoeDetail.Stock, &name, &price)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Shoe detail not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch shoe detail"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"shoe_id":  shoeDetail.ShoeID,
		"model_id": shoeDetail.ModelID,
		"size":     shoeDetail.Size,
		"stock":    shoeDetail.Stock,
		"name":     name,
		"price":    price,
	})
}

// CreateShoeDetailRequest represents the request body for creating a shoe detail and model
type CreateShoeDetailRequest struct {
	ShoeDetail models.ShoeDetail `json:"shoe_detail"`
	ShoeModel  models.ShoeModel  `json:"shoe_model"`
}

// CreateShoeDetail godoc
// @Summary      Create a new shoe detail and model
// @Description  Add a new shoe detail and model to the database
// @Tags         shoe-details
// @Accept       json
// @Produce      json
// @Param        request body      CreateShoeDetailRequest true "Shoe Detail and Model"
// @Success      201       {object}  models.ShoeDetail
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /shoe-details [post]
func CreateShoeDetail(c echo.Context) error {
	request := new(CreateShoeDetailRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	// Insert the shoe model first
	modelQuery := "INSERT INTO shoe_models (name, price) VALUES (?, ?)"
	modelRes, err := config.DB.Exec(modelQuery, request.ShoeModel.Name, request.ShoeModel.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to insert shoe model"})
	}

	// Get the last inserted model ID
	lastModelID, err := modelRes.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve last insert model ID"})
	}

	// Now insert the shoe detail with the new model ID
	shoeDetail := request.ShoeDetail
	shoeDetail.ModelID = int(lastModelID) // Set the model ID to the newly inserted model ID

	query := "INSERT INTO shoe_details (model_id, size, stock) VALUES (?, ?, ?)"
	res, err := config.DB.Exec(query, shoeDetail.ModelID, shoeDetail.Size, shoeDetail.Stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to insert shoe detail"})
	}

	// Get the last inserted shoe detail ID
	lastShoeDetailID, err := res.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve last insert shoe detail ID"})
	}

	shoeDetail.ShoeID = int(lastShoeDetailID) // Set the shoe ID to the newly inserted shoe detail ID
	return c.JSON(http.StatusCreated, shoeDetail)
}

type UpdateShoeDetailRequest struct {
	ShoeDetail models.ShoeDetail `json:"shoe_detail"`
	ShoeModel  models.ShoeModel  `json:"shoe_model"`
}

// UpdateShoeDetail godoc
// @Summary      Update an existing shoe detail by ID
// @Description  Modify an existing shoe detail and shoe model using the shoe ID
// @Tags         shoe-details
// @Accept       json
// @Produce      json
// @Param        id path int true "Shoe ID"
// @Param        updateRequest body UpdateShoeDetailRequest true "Update Shoe Detail Request"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /shoe-details/{id} [put]

func UpdateShoeDetail(c echo.Context) error {
	shoeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid shoe ID"})
	}

	updateRequest := new(UpdateShoeDetailRequest)
	if err := c.Bind(updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	// Update Shoe Detail
	shoeDetail := updateRequest.ShoeDetail
	query := "UPDATE shoe_details SET model_id = ?, size = ?, stock = ? WHERE shoe_id = ?"
	_, err = config.DB.Exec(query, shoeDetail.ModelID, shoeDetail.Size, shoeDetail.Stock, shoeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update shoe detail"})
	}

	// Update Shoe Model
	shoeModel := updateRequest.ShoeModel
	modelQuery := "UPDATE shoe_models SET name = ?, price = ? WHERE model_id = (SELECT model_id FROM shoe_details WHERE shoe_id = ?)"
	_, err = config.DB.Exec(modelQuery, shoeModel.Name, shoeModel.Price, shoeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update shoe model"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Shoe detail updated successfully"})
}

// DeleteShoeDetail godoc
// @Summary      Delete a shoe detail by ID
// @Description  Remove a shoe detail from the database using the shoe ID
// @Tags         shoe-details
// @Accept       json
// @Produce      json
// @Param        id path int true "Shoe ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /shoe-details/{id} [delete]
func DeleteShoeDetail(c echo.Context) error {
	shoeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid shoe ID"})
	}

	query := "DELETE FROM shoe_details WHERE shoe_id = ?"
	_, err = config.DB.Exec(query, shoeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete shoe detail"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Shoe detail deleted successfully"})
}
