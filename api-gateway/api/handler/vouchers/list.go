package vouchers

import (
	VOUCHERS_PRESENTATION "api-gateway/api/presentation/vouchers"
	"api-gateway/entity/responses"
	"api-gateway/entity/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler *Controller) List(c echo.Context) error {
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

	voucherList, err := handler.vs.GetVouchers()
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	voucherListResponse := make([]VOUCHERS_PRESENTATION.VoucherResponse, len(voucherList))
	for i, voucher := range voucherList {
		voucherListResponse[i] = VOUCHERS_PRESENTATION.VoucherResponse{
			VoucherID:  voucher.VoucherID,
			Discount:   voucher.Discount,
			ValidUntil: voucher.ValidUntil,
			Used:       voucher.Used,
		}
	}
	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    voucherListResponse,
	}
	return responses.HandleResponse(c, res)
}
