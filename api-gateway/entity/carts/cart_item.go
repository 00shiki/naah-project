package carts

type CartItem struct {
	CartID   int32 `json:"cart_id"`
	UserID   int32 `json:"user_id"`
	ShoeID   int32 `json:"shoe_id"`
	Quantity int32 `json:"quantity"`
}
