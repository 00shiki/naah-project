package vouchers

import (
	VOUCHERS_PRESENTATION "api-gateway/api/presentation/vouchers"
	"api-gateway/entity/responses"
	"api-gateway/entity/users"
	"api-gateway/entity/vouchers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) Create(c echo.Context) error {
	role, ok := c.Get("role").(float64)
	if !ok {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Token",
		}
		return responses.HandleResponse(c, res)
	}

	if users.Role(role) != users.Admin {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Restricted Access",
		}
		return responses.HandleResponse(c, res)
	}

	payload := new(VOUCHERS_PRESENTATION.CreateRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	voucher := &vouchers.Voucher{
		VoucherID:  payload.VoucherID,
		Discount:   payload.Discount,
		ValidUntil: payload.ValidUntil,
	}
	err := handler.vs.CreateVoucher(voucher)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	res := &responses.Response{
		Code:    http.StatusCreated,
		Message: "Voucher Created",
	}
	return responses.HandleResponse(c, res)
}
