package vouchers

import (
	"api-gateway/entity/responses"
	"api-gateway/entity/users"
	"net/http"

	VOUCHERS_PRESENTATION "api-gateway/api/presentation/vouchers"

	"github.com/labstack/echo/v4"
)

func (handler *Controller) Detail(c echo.Context) error {
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

	voucherID := c.Param("voucherID")

	voucher, err := handler.vs.GetVoucherByID(voucherID)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data: VOUCHERS_PRESENTATION.VoucherResponse{
			VoucherID:  voucher.VoucherID,
			Discount:   voucher.Discount,
			ValidUntil: voucher.ValidUntil,
			Used:       voucher.Used,
		},
	}
	return responses.HandleResponse(c, res)
}
