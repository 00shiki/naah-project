package handler

import (
	"context"
	"database/sql"
	"order-service/pb"
)

type OrderHandler struct {
	db *sql.DB
}

// type OrderItem struct {
// }

func NewOrderHandler(db *sql.DB) *OrderHandler {

	return &OrderHandler{db}
}

func (h *CartHandler) AddOrder(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartResponse, error) {
	return nil, nil
}
