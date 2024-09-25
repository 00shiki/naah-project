package carts

import (
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

func (handler *Controller) Delete(c echo.Context) error {
	cartIDStr := c.Param("cartID")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	err = handler.cs.DeleteCartItem(int32(cartID))
	if err != nil {
		res := &responses.Response{
			Message: err.Error(),
		}
		switch status.Code(err) {
		case codes.NotFound:
			res.Code = http.StatusNotFound
		case codes.Internal:
			res.Code = http.StatusInternalServerError
		}
		return responses.HandleResponse(c, res)
	}

	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
	}
	return responses.HandleResponse(c, res)
}
