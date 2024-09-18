package handler

import (
	"context"
	"database/sql"
	"log"
	"order-service/pb"
)

type UserHandler struct {
	db *sql.DB
}

type CartItem struct {
	cart_id  int32
	user_id  int32
	shoes_id int32
	quantity int32
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db}
}

func (h *UserHandler) AddCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartResponse, error) {
	// Debugging: Log successful response

	item := pb.CartItem{
		CartId:   1,
		UserId:   1,
		ShoesId:  1,
		Quantity: 1,
	}
	response := &pb.AddCartResponse{Item: &item}
	log.Printf("Response: %+v", response)

	return response, nil
}

func (h *UserHandler) SubstractCart(ctx context.Context, req *pb.SubstractCartRequest) (*pb.SubstractCartResponse, error) {
	// Debugging: Log successful response

	item := pb.CartItem{
		CartId:   1,
		UserId:   1,
		ShoesId:  1,
		Quantity: 1,
	}
	response := &pb.SubstractCartResponse{Item: &item}
	log.Printf("Response: %+v", response)

	return response, nil
}

func (h *UserHandler) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	// Create a CartItem
	item := pb.CartItem{
		CartId:   1,
		UserId:   1,
		ShoesId:  1,
		Quantity: 1,
	}

	// Create a slice of CartItem (use [] to create a slice)
	items := []*pb.CartItem{&item}

	// Create the GetCartResponse and assign the slice of CartItems
	response := &pb.GetCartResponse{
		Items: items, // Assign the slice directly to the 'Items' field
	}

	log.Printf("Response: %+v", response)

	return response, nil
}
