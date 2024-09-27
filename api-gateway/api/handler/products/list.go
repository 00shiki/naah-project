package products

import (
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) List(c echo.Context) error {
	shoes, err := handler.ps.GetAllShoes()
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
		Data:    shoes,
	}
	return responses.HandleResponse(c, res)
}
