package deliveries

import "api-gateway/entity/carts"

type Delivery struct {
	ID                int32
	OrderID           int32
	TrackID           int32
	OriginCityID      string
	DestinationCityID string
	Courier           Courier
	Carts             []carts.CartItem
}

type Courier struct {
	Name    string
	Service string
}
