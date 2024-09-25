package handlers

import (
	"net/http"
	"product-service/config"
	"product-service/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// admin create
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

// admin retrieve
func GetShoeModels(c echo.Context) error {
	var shoeModels []models.ShoeModel
	if err := config.DB.Find(&shoeModels).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch shoe models", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, shoeModels)
}

// admin retrieve by id
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

// admin update
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

// admin delete
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
