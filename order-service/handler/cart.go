package handler

import (
	"context"
	"database/sql"
	"log"
	"order-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartHandler struct {
	db *sql.DB
}

// type CartItem struct {
// 	cart_id  int32
// 	user_id  int32
// 	shoes_id int32
// 	quantity int32
// }

func NewCartHandler(db *sql.DB) *CartHandler {
	return &CartHandler{db}
}

func (h *CartHandler) AddCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartResponse, error) {
	var cartId int
	var currentQuantity int

	// Check if the cart already exists for the given user_id and shoe_id
	query := "SELECT cart_id, quantity FROM carts WHERE user_id = ? AND shoe_id = ?"
	err := h.db.QueryRow(query, req.UserId, req.ShoeId).Scan(&cartId, &currentQuantity)
	if err != nil {
		if err == sql.ErrNoRows {
			// No existing cart entry for the user and shoe combination
			cartId = 0
			currentQuantity = 0
			log.Println("No cart found for user and shoe.")
		} else {
			// Handle unexpected database errors and return a gRPC Internal error
			log.Println("Error querying database:", err)
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}
	}

	if cartId == 0 {
		// No existing cart, insert a new cart entry
		query = "INSERT INTO carts (user_id, shoe_id, quantity) VALUES (?, ?, ?)"
		_, err := h.db.Exec(query, req.UserId, req.ShoeId, req.Quantity)
		if err != nil {
			// Handle insertion error and return a gRPC Internal error
			log.Println("Error inserting into database:", err)
			return nil, status.Errorf(codes.Internal, "Error inserting into database: %v", err)
		}
		log.Println("New cart item added to the database.")
		currentQuantity = int(req.Quantity) // Since it's a new entry, the quantity is the one added
	} else {
		// Cart exists, update the quantity
		newQuantity := currentQuantity + int(req.Quantity)
		query = "UPDATE carts SET quantity = ? WHERE cart_id = ?"
		_, err := h.db.Exec(query, newQuantity, cartId)
		if err != nil {
			// Handle update error and return a gRPC Internal error
			log.Println("Error updating database:", err)
			return nil, status.Errorf(codes.Internal, "Error updating database: %v", err)
		}
		log.Println("Cart item updated in the database.")
		currentQuantity = newQuantity // Updated quantity after the increase
	}

	// Construct the response with the cart item details, including the final quantity
	item := pb.CartItem{
		CartId:   int32(cartId),
		UserId:   req.UserId,
		ShoeId:   req.ShoeId,
		Quantity: int32(currentQuantity), // Return the final quantity after the operation
	}

	response := &pb.AddCartResponse{
		Item: &item,
	}

	// Log and return the response
	log.Printf("Response: %+v", response)
	return response, nil
}

func (h *CartHandler) SubtractCart(ctx context.Context, req *pb.SubtractCartRequest) (*pb.SubtractCartResponse, error) {
	var cartId int
	var currentQuantity int

	// Check if the cart already exists for the given user_id and shoe_id
	query := "SELECT cart_id, quantity FROM carts WHERE user_id = ? AND shoe_id = ?"
	err := h.db.QueryRow(query, req.UserId, req.ShoeId).Scan(&cartId, &currentQuantity)
	if err != nil {
		if err == sql.ErrNoRows {
			// No existing cart entry for the user and shoe combination
			log.Println("No cart found for user and shoe. Cannot subtract.")
			return nil, status.Errorf(codes.NotFound, "No cart found for user_id %d and shoe_id %d", req.UserId, req.ShoeId)
		} else {
			// Return internal gRPC error for database issues
			log.Println("Error querying database:", err)
			return nil, status.Errorf(codes.Internal, "Database error: %v", err)
		}
	}

	// Check if the current quantity is enough for subtraction
	newQuantity := currentQuantity - int(req.Quantity)
	if newQuantity <= 0 {
		// If the result is zero or negative, delete the cart item
		query = "DELETE FROM carts WHERE cart_id = ?"
		_, err := h.db.Exec(query, cartId)
		if err != nil {
			// Return internal gRPC error for database issues during deletion
			log.Println("Error deleting from database:", err)
			return nil, status.Errorf(codes.Internal, "Error deleting cart item: %v", err)
		}
		log.Println("Cart item deleted from the database.")
		newQuantity = 0 // Set newQuantity to 0 after deletion
	} else {
		// Otherwise, update the quantity
		query = "UPDATE carts SET quantity = ? WHERE cart_id = ?"
		_, err := h.db.Exec(query, newQuantity, cartId)
		if err != nil {
			// Return internal gRPC error for database issues during update
			log.Println("Error updating database:", err)
			return nil, status.Errorf(codes.Internal, "Error updating cart item: %v", err)
		}
		log.Println("Cart item updated in the database.")
	}

	// Construct the response with the updated cart item details
	item := pb.CartItem{
		CartId:   int32(cartId),
		UserId:   req.UserId,
		ShoeId:   req.ShoeId,
		Quantity: int32(newQuantity), // Return the new quantity after subtraction or deletion
	}

	response := &pb.SubtractCartResponse{
		Item: &item,
	}

	// Log and return the response
	log.Printf("Response: %+v", response)
	return response, nil
}

func (h *CartHandler) GetCartsByUserId(ctx context.Context, req *pb.GetCartsByUserIdRequest) (*pb.GetCartsByUserIdResponse, error) {
	query := "SELECT cart_id, user_id, shoe_id, quantity FROM carts WHERE user_id = ?"
	rows, err := h.db.Query(query, req.UserId)
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
	}
	defer rows.Close()

	var carts []*pb.CartItem
	for rows.Next() {
		var cartItem pb.CartItem
		if err := rows.Scan(&cartItem.CartId, &cartItem.UserId, &cartItem.ShoeId, &cartItem.Quantity); err != nil {
			log.Println("Error scanning row:", err)
			return nil, status.Errorf(codes.Internal, "Error scanning row: %v", err)
		}
		carts = append(carts, &cartItem)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, status.Errorf(codes.Internal, "Error with rows: %v", err)
	}

	response := &pb.GetCartsByUserIdResponse{
		Carts: carts,
	}
	return response, nil
}

func (h *CartHandler) DeleteCartByCartId(ctx context.Context, req *pb.DeleteCartByCartIdRequest) (*pb.DeleteCartByCartIdResponse, error) {
	query := "DELETE FROM carts WHERE cart_id = ?"
	result, err := h.db.Exec(query, req.CartId)
	if err != nil {
		log.Println("Error deleting from database:", err)
		return nil, status.Errorf(codes.Internal, "Error deleting from database: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return nil, status.Errorf(codes.Internal, "Error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		log.Println("No cart found with the given cart_id.")
		return nil, status.Errorf(codes.NotFound, "No cart found with cart_id %d", req.CartId)
	}

	response := &pb.DeleteCartByCartIdResponse{
		Message: "Cart deleted successfully.",
	}
	return response, nil
}
