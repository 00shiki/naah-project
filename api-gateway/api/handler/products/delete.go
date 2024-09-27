package products

import (
	"api-gateway/entity/responses"
	"api-gateway/entity/users"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (handler *Controller) Delete(c echo.Context) error {
	role, ok := c.Get("role").(float64)
	if !ok {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Token",
		}
		return responses.HandleResponse(c, res)
	}

	if users.Role(role) != users.Admin {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Restricted Access",
		}
		return responses.HandleResponse(c, res)
	}

	productIDStr := c.Param("productID")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	err = handler.ps.DeleteProduct(int32(productID))
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
	}
	return responses.HandleResponse(c, res)
}
