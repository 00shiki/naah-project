package vouchers

import "api-gateway/entity/vouchers"

type Service interface {
	CreateVoucher(voucher *vouchers.Voucher) error
	GetVoucherByID(voucherID string) (*vouchers.Voucher, error)
	GetVouchers() ([]vouchers.Voucher, error)
}
