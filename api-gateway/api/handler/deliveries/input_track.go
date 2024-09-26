package deliveries

import (
	DELIVERIES_PRESENTATION "api-gateway/api/presentation/deliveries"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) InputTrack(c echo.Context) error {
	payload := new(DELIVERIES_PRESENTATION.InputTrackRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	err := handler.ds.InputTrackID(payload.OrderID, payload.TrackID)
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
