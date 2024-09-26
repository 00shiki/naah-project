package deliveries

import (
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) ListCourier(c echo.Context) error {
	couriers, err := handler.ds.GetCouriers()
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
		Data:    couriers,
	}
	return responses.HandleResponse(c, res)
}

func (handler *Controller) ListProvince(c echo.Context) error {
	provinces, err := handler.ds.GetProvinces()
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
		Data:    provinces,
	}
	return responses.HandleResponse(c, res)
}

func (handler *Controller) ListCity(c echo.Context) error {
	provinceID := c.Param("provinceID")

	cities, err := handler.ds.GetCities(provinceID)
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
		Data:    cities,
	}
	return responses.HandleResponse(c, res)
}
