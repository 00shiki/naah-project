package orders

import (
	"api-gateway/entity/deliveries"
	"api-gateway/entity/products"
)

type Order struct {
	ID          int32
	UserID      int32
	VoucherID   string
	Status      string
	Price       int32
	Fee         int32
	Discount    float64
	TotalPrice  int32
	Delivery    deliveries.Delivery
	Metadata    string
	OrderItems  []products.ShoeDetail
	InvoiceUrl  string
	ExpiredDate string
}
