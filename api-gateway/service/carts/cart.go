package carts

import (
	"api-gateway/entity/carts"
	pb "api-gateway/proto"
	"context"
	"time"
)

type CartService struct {
	client pb.CartServiceClient
}

func NewCartService(client pb.CartServiceClient) *CartService {
	return &CartService{
		client: client,
	}
}

func (cs *CartService) UpdateCart(cart *carts.CartItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.UpdateCartRequest{
		UserId:   cart.UserID,
		ShoeId:   cart.ShoeID,
		Quantity: cart.Quantity,
	}
	res, err := cs.client.UpdateCart(ctx, req)
	if err != nil {
		return err
	}
	cart.CartID = res.GetItem().GetCartId()
	cart.UserID = res.GetItem().GetUserId()
	cart.ShoeID = res.GetItem().GetShoeId()
	cart.Quantity = res.GetItem().GetQuantity()
	return nil
}

func (cs *CartService) GetUserCarts(userID int64) ([]*carts.CartItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.GetCartsByUserIdRequest{
		UserId: int32(userID),
	}
	res, err := cs.client.GetCartsByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	userCarts := make([]*carts.CartItem, len(res.GetCarts()))
	for i, c := range res.GetCarts() {
		userCarts[i] = &carts.CartItem{
			CartID:   c.CartId,
			UserID:   c.UserId,
			ShoeID:   c.ShoeId,
			Quantity: c.Quantity,
		}
	}
	return userCarts, nil
}

func (cs *CartService) DeleteCartItem(cartID int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.DeleteCartByCartIdRequest{
		CartId: cartID,
	}
	_, err := cs.client.DeleteCartByCartId(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
