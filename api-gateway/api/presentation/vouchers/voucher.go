package vouchers

type VoucherResponse struct {
	VoucherID  string  `json:"voucher_id"`
	Discount   float64 `json:"discount"`
	ValidUntil string  `json:"valid_until"`
	Used       bool    `json:"used"`
}
