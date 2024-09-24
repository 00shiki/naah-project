package carts

import "api-gateway/entity/carts"

type Service interface {
	UpdateCart(cart *carts.CartItem) error
	GetUserCarts(userID int64) ([]*carts.CartItem, error)
	DeleteCartItem(cartID int32) error
}
