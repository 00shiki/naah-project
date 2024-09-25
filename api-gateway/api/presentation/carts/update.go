package carts

type UpdateRequest struct {
	ShoeID   int32 `json:"shoe_id" validate:"required"`
	Quantity int32 `json:"quantity" validate:"required"`
}

type UpdateResponse struct {
	CartID   int32 `json:"cart_id"`
	ShoeID   int32 `json:"shoe_id"`
	Quantity int32 `json:"quantity"`
}
