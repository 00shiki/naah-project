package deliveries

import (
	DELIVERIES_PRESENTATION "api-gateway/api/presentation/deliveries"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) Callback(c echo.Context) error {
	payload := new(DELIVERIES_PRESENTATION.CallbackRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	err := handler.ds.CallbackDelivery(payload.TrackID, payload.Status)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	return c.NoContent(http.StatusNoContent)
}
