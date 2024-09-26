package vouchers

type Voucher struct {
	VoucherID  string
	Discount   float64
	ValidUntil string
	Used       bool
}
