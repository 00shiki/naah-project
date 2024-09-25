package orders

import (
	ORDERS_PRESENTATION "api-gateway/api/presentation/orders"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) Callback(c echo.Context) error {
	payload := new(ORDERS_PRESENTATION.CallbackRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	err := handler.os.CallbackNotification(payload.ExternalID, payload.Status, payload.PaidAmount)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	return c.NoContent(http.StatusOK)
}
