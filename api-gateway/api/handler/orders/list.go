package orders

import (
	ORDERS_PRESENTATION "api-gateway/api/presentation/orders"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (handler *Controller) List(c echo.Context) error {
	userID, ok := c.Get("user_id").(float64)
	if !ok {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Token",
		}
		return responses.HandleResponse(c, res)
	}

	userOrders, err := handler.os.UserOrders(int32(userID))
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

	listResponses := make([]ORDERS_PRESENTATION.ListResponse, len(userOrders))
	for i, order := range userOrders {
		orderItems := make([]ORDERS_PRESENTATION.OrderItemResponse, len(order.OrderItems))
		for j, orderItem := range order.OrderItems {
			orderItems[j] = ORDERS_PRESENTATION.OrderItemResponse{
				Name:     orderItem.Name,
				Price:    orderItem.Price,
				Quantity: orderItem.Stock,
				Size:     orderItem.Size,
			}
		}
		listResponses[i] = ORDERS_PRESENTATION.ListResponse{
			OrderID:    order.ID,
			Status:     order.Status,
			Fee:        order.Fee,
			OrderItems: orderItems,
			Discount:   order.Discount,
			TotalPrice: order.TotalPrice,
			VoucherID:  order.VoucherID,
		}
	}
	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    listResponses,
	}
	return responses.HandleResponse(c, res)

}
