package handlers

import (
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

	if err := config.DB.Create(&shoeModel).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create shoe model", "error": err.Error()})
	}

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
	var shoeModels []models.ShoeModel
	if err := config.DB.Find(&shoeModels).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch shoe models", "error": err.Error()})
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

	var shoeModel models.ShoeModel
	if err := config.DB.First(&shoeModel, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Shoe model not found"})
	}
	return c.JSON(http.StatusOK, shoeModel)
}

// UpdateShoeModel godoc
// @Summary Update a shoe model
// @Description Update the details of a shoe model by ID
// @Tags shoe-models
// @Accept json
// @Produce json
// @Param id path int true "Shoe Model ID"
// @Param shoeModel body models.ShoeModel true "Shoe Model"
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

	var shoeModel models.ShoeModel
	if err := config.DB.First(&shoeModel, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Shoe model not found"})
	}

	if err := c.Bind(&shoeModel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input", "error": err.Error()})
	}

	if err := config.DB.Save(&shoeModel).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not update shoe model", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, shoeModel)
}

// DeleteShoeModel godoc
// @Summary Delete a shoe model
// @Description Delete a shoe model by ID
// @Tags shoe-models
// @Accept json
// @Produce json
// @Param id path int true "Shoe Model ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/shoe-models/{id} [delete]
func DeleteShoeModel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid shoe model ID"})
	}

	var shoeModel models.ShoeModel
	if err := config.DB.First(&shoeModel, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Shoe model not found"})
	}

	if err := config.DB.Delete(&shoeModel).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not delete shoe model", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Shoe model deleted successfully"})
}
