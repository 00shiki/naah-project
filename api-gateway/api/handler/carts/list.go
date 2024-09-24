package carts

import (
	CARTS_PRESENTATION "api-gateway/api/presentation/carts"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) ListUserCarts(c echo.Context) error {
	userID, ok := c.Get("user_id").(float64)
	if !ok {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Token",
		}
		return responses.HandleResponse(c, res)
	}

	userCarts, err := handler.cs.GetUserCarts(int64(userID))
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	userCartsRes := make([]CARTS_PRESENTATION.ListUserCartsResponse, len(userCarts))
	for i, userCart := range userCarts {
		userCartsRes[i] = CARTS_PRESENTATION.ListUserCartsResponse{
			CartID:   userCart.CartID,
			ShoeID:   userCart.ShoeID,
			Quantity: userCart.Quantity,
		}
	}
	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userCartsRes,
	}
	return responses.HandleResponse(c, res)
}
