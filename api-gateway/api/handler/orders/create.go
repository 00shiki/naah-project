package orders

import (
	ORDERS_PRESENTATION "api-gateway/api/presentation/orders"
	"api-gateway/entity/deliveries"
	"api-gateway/entity/orders"
	"api-gateway/entity/products"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) Create(c echo.Context) error {
	userID, ok := c.Get("user_id").(float64)
	if !ok {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Token",
		}
		return responses.HandleResponse(c, res)
	}

	payload := new(ORDERS_PRESENTATION.CreateRequest)
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

	carts := make([]products.Product, len(payload.CartIDs))
	for i, cartID := range payload.CartIDs {
		carts[i] = products.Product{
			ID: cartID,
		}
	}
	order := &orders.Order{
		UserID:     int32(userID),
		VoucherID:  payload.VoucherID,
		OrderItems: carts,
		Delivery: deliveries.Delivery{
			Courier: deliveries.Courier{
				Name:    payload.CourierName,
				Service: payload.CourierServiceName,
			},
			OriginCityID:      payload.OriginCityID,
			DestinationCityID: payload.DestinationCityID,
		},
		Fee:      payload.OtherFee,
		Metadata: payload.Metadata,
	}
	err := handler.os.CreateOrder(order)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	res := &responses.Response{
		Code:    http.StatusCreated,
		Message: "Success",
		Data: ORDERS_PRESENTATION.CreateResponse{
			InvoiceUrl:  order.InvoiceUrl,
			ExpiredDate: order.ExpiredDate,
			TotalPrice:  order.TotalPrice,
		},
	}
	return responses.HandleResponse(c, res)
}
