package orders

type ListResponse struct {
	OrderID    int32               `json:"order_id"`
	Status     string              `json:"status"`
	OrderItems []OrderItemResponse `json:"order_items"`
	Fee        int32               `json:"fee"`
	Discount   float64             `json:"discount"`
	TotalPrice int32               `json:"total_price"`
	VoucherID  string              `json:"voucher_id"`
}

type OrderItemResponse struct {
	Name     string `json:"name"`
	Price    int32  `json:"price"`
	Quantity int32  `json:"quantity"`
	Size     int32  `json:"size"`
}
