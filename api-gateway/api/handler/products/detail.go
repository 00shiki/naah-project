package products

import (
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (handler *Controller) Detail(c echo.Context) error {
	productIDStr := c.Param("productID")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	product, err := handler.ps.GetShoeModelByID(int32(productID))
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    product,
	}
	return responses.HandleResponse(c, res)
}
