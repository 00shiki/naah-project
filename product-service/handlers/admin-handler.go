package handlers

import (
	"database/sql"
	"net/http"
	"product-service/config"
	"product-service/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateShoeModel godoc
// @Summary Create a new shoe model
// @Description Create a new shoe model with the given data
// @Tags shoe-models
// @Accept json
// @Produce json
// @Param shoeModel body models.ShoeModel true "Shoe Model"
// @Success 201 {object} models.ShoeModel
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/shoe-models [post]
func CreateShoeModel(c echo.Context) error {
	shoeModel := new(models.ShoeModel)
	if err := c.Bind(shoeModel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input", "error": err.Error()})
	}

	query := "INSERT INTO shoe_models (name, price) VALUES (?, ?)"
	result, err := config.DB.Exec(query, shoeModel.Name, shoeModel.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create shoe model", "error": err.Error()})
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not retrieve inserted ID", "error": err.Error()})
	}

	shoeModel.ModelID = int(id)
	return c.JSON(http.StatusCreated, shoeModel)
}

// GetShoeModels godoc
// @Summary Retrieve all shoe models
// @Description Get a list of all shoe models
// @Tags shoe-models
// @Accept json
// @Produce json
// @Success 200 {array} models.ShoeModel
// @Failure 500 {object} map[string]string
// @Router /admin/shoe-models [get]
func GetShoeModels(c echo.Context) error {
	query := "SELECT model_id, name, price FROM shoe_models"
	rows, err := config.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch shoe models", "error": err.Error()})
	}
	defer rows.Close()

	var shoeModels []models.ShoeModel
	for rows.Next() {
		var shoeModel models.ShoeModel
		if err := rows.Scan(&shoeModel.ModelID, &shoeModel.Name, &shoeModel.Price); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not scan shoe model", "error": err.Error()})
		}
		shoeModels = append(shoeModels, shoeModel)
	}
	return c.JSON(http.StatusOK, shoeModels)
}

// GetShoeModelByID godoc
// @Summary Retrieve a shoe model by ID
// @Description Get a shoe model by its ID
// @Tags shoe-models
// @Accept json
// @Produce json
// @Param id path int true "Shoe Model ID"
// @Success 200 {object} models.ShoeModel
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /admin/shoe-models/{id} [get]
func GetShoeModelByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid shoe model ID"})
	}

	query := "SELECT model_id, name, price FROM shoe_models WHERE model_id = ?"
	row := config.DB.QueryRow(query, id)

	var shoeModel models.ShoeModel
	if err := row.Scan(&shoeModel.ModelID, &shoeModel.Name, &shoeModel.Price); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Shoe model not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch shoe model", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, shoeModel)
}

// UpdateShoeModel godoc
// @Summary Update a shoe model
// @Description Update the shoe model with the given ID
// @Tags shoe-models
// @Accept json
// @Produce json
// @Param id path int true "Shoe Model ID"
// @Param shoeModel body models.ShoeModel true "Updated Shoe Model"
// @Success 200 {object} models.ShoeModel
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/shoe-models/{id} [put]
func UpdateShoeModel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid shoe model ID"})
	}

	shoeModel := new(models.ShoeModel)
	if err := c.Bind(shoeModel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input", "error": err.Error()})
	}

	query := "UPDATE shoe_models SET name = ?, price = ? WHERE model_id = ?"
	_, err = config.DB.Exec(query, shoeModel.Name, shoeModel.Price, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not update shoe model", "error": err.Error()})
	}

	shoeModel.ModelID = id
	return c.JSON(http.StatusOK, shoeModel)
}

func CreateShoeDetail(c echo.Context) error {
	modelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid shoe model ID"})
	}

	payload := new(models.ShoeDetail)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input", "error": err.Error()})
	}
	query := "INSERT INTO shoe_details (size, stock, model_id) VALUES (?, ?, ?)"
	result, err := config.DB.Exec(query, payload.Size, payload.Stock, modelID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create shoe detail", "error": err.Error()})
	}
	id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not retrieve inserted ID", "error": err.Error()})
	}
	payload.ShoeID = int(id)
	return c.JSON(http.StatusCreated, payload)
}

// DeleteShoeModel godoc
// @Summary Delete a shoe model
// @Description Delete the shoe model with the given ID
// @Tags shoe-models
// @Accept json
// @Produce json
// @Param id path int true "Shoe Model ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/shoe-models/{id} [delete]
func DeleteShoeModel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid shoe model ID"})
	}

	query := "DELETE FROM shoe_models WHERE model_id = ?"
	result, err := config.DB.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not delete shoe model", "error": err.Error()})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not check if shoe model was deleted", "error": err.Error()})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Shoe model not found"})
	}

	return c.NoContent(http.StatusNoContent)
}
