package carts

type ListUserCartsResponse struct {
	CartID   int32 `json:"cart_id"`
	ShoeID   int32 `json:"shoe_id"`
	Quantity int32 `json:"quantity"`
}
