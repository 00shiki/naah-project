package orders

type CreateRequest struct {
	VoucherID          string  `json:"voucher_id"`
	CartIDs            []int32 `json:"cart_ids"`
	CourierName        string  `json:"courier_name"`
	CourierServiceName string  `json:"courier_service_name"`
	OriginCityID       string  `json:"origin_city_id"`
	DestinationCityID  string  `json:"destination_city_id"`
	OtherFee           int32   `json:"other_fee"`
	Metadata           string  `json:"metadata"`
}

type CreateResponse struct {
	InvoiceUrl  string
	ExpiredDate string
	TotalPrice  int32
}
