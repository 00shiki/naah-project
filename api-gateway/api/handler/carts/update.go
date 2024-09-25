package carts

import (
	CARTS_PRESENTATION "api-gateway/api/presentation/carts"
	"api-gateway/entity/carts"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (handler *Controller) Update(c echo.Context) error {
	userID, ok := c.Get("user_id").(float64)
	if !ok {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Token",
		}
		return responses.HandleResponse(c, res)
	}

	payload := new(CARTS_PRESENTATION.UpdateRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}
	if err := c.Validate(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	cart := &carts.CartItem{
		UserID:   int32(userID),
		ShoeID:   payload.ShoeID,
		Quantity: payload.Quantity,
	}

	err := handler.cs.UpdateCart(cart)
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
		Data: CARTS_PRESENTATION.UpdateResponse{
			CartID:   cart.CartID,
			ShoeID:   cart.ShoeID,
			Quantity: cart.Quantity,
		},
	}
	return responses.HandleResponse(c, res)
}
