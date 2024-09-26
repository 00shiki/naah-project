package deliveries

type DeliveryCostRequest struct {
	OriginID      string  `json:"origin_id"`
	DestinationID string  `json:"destination_id"`
	CartIDs       []int32 `json:"cart_ids"`
	Courier       string  `json:"courier"`
}

type DeliveryCostResponse struct {
	Origin      DeliveryItemResponse  `json:"origin"`
	Destination DeliveryItemResponse  `json:"destination"`
	Services    []ServiceItemResponse `json:"services"`
}

type DeliveryItemResponse struct {
	CityID       string `json:"city_id"`
	CityName     string `json:"city_name"`
	ProvinceID   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
	Type         string `json:"type"`
	PostalCode   string `json:"postal_code"`
}

type ServiceItemResponse struct {
	ServiceName string `json:"service_name"`
	Description string `json:"description"`
	Cost        int32  `json:"cost"`
	Etd         string `json:"etd"`
}
