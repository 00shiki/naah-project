package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log" // Added for logging
	"order-service/pb"
	"order-service/service"
	"order-service/utils"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderHandler struct {
	db *sql.DB
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{db}
}

type shoes struct {
	ID    int
	Price int
	Qty   int
}

func (h *OrderHandler) AddOrder(ctx context.Context, req *pb.AddOrderRequest) (*pb.AddOrderResponse, error) {
	log.Println("Starting AddOrder function")

	// Start a transaction
	tx, err := h.db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Error starting transaction: %v", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("Error rolling back transaction: %v\n", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				log.Printf("Error committing transaction: %v\n", cmErr)
			}
		}
	}()

	cartIds := utils.RemoveDuplicates(req.CartIds)
	log.Printf("Unique cart IDs: %v\n", cartIds)

	var shoeId, shoeQty, shoePrice int
	var shoeList []shoes
	var userIdTemp int

	// Fetching shoes from cart
	for _, cartId := range cartIds {
		query := "SELECT user_id, shoe_id, quantity FROM carts WHERE cart_id = ?"
		err := tx.QueryRow(query, cartId).Scan(&userIdTemp, &shoeId, &shoeQty)
		if err != nil {
			log.Printf("Error querying database for cart_id %d: %v\n", cartId, err)
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}
		log.Printf("Cart ID: %d, Shoe ID: %d, Quantity: %d\n", cartId, shoeId, shoeQty)

		if userIdTemp != int(req.UserId) {
			return nil, status.Errorf(codes.InvalidArgument, "cart is not belong to user")
		}

		query = `SELECT sm.price FROM shoe_models sm JOIN shoe_details sd ON sm.model_id = sd.model_id WHERE sd.shoe_id = ?`
		err = tx.QueryRow(query, shoeId).Scan(&shoePrice)
		if err != nil {
			log.Printf("Error querying database for shoe_id %d: %v\n", shoeId, err)
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}
		log.Printf("Shoe ID: %d, Price: %d\n", shoeId, shoePrice)

		shoeList = append(shoeList, shoes{ID: shoeId, Price: shoePrice, Qty: shoeQty})
	}

	// Calculate total price
	var price int
	for _, shoe := range shoeList {
		price += shoe.Price * shoe.Qty
	}
	log.Printf("Calculated total price from cart: %d\n", price)

	// Voucher processing
	var voucherID string
	if req.VoucherId == "" {
		voucherID = "NOVOUCHER"
	} else {
		voucherID = req.VoucherId
	}
	log.Printf("Using voucher ID: %s\n", voucherID)

	var discountPercent float64
	var validUntilStr string
	var used bool

	query := "SELECT discount, valid_until, used FROM vouchers WHERE voucher_id = ?"
	log.Printf("Executing query: %s\n", query)
	err = tx.QueryRow(query, voucherID).Scan(&discountPercent, &validUntilStr, &used)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No voucher found for VoucherId=%s\n", voucherID)
			return nil, status.Errorf(codes.NotFound, "Voucher not found")
		} else {
			log.Printf("Error querying database for VoucherId=%s: %v\n", voucherID, err)
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}
	}

	layout := "2006-01-02"
	validUntil, err := time.Parse(layout, validUntilStr)
	if err != nil {
		log.Printf("Error parsing validUntil: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Error parsing validUntil: %v", err)
	}
	log.Printf("Voucher valid until: %s\n", validUntil)

	if validUntil.Before(time.Now()) {
		log.Println("Voucher has expired")
		return nil, status.Errorf(codes.InvalidArgument, "Voucher has expired")
	} else if used {
		log.Println("Voucher has already been used")
		return nil, status.Errorf(codes.InvalidArgument, "Voucher has already been used")
	}

	discount := int(float64(price) * (discountPercent / 100))
	log.Printf("Applying discount: %d on price: %d\n", discount, price)

	// Calculate total weight
	var weight int32
	for _, cartID := range req.CartIds {
		query := "SELECT quantity FROM carts WHERE cart_id = ?"
		log.Printf("Running query: %s with cart_id: %d", query, cartID)

		var qty int32
		err := tx.QueryRow(query, cartID).Scan(&qty)
		if err != nil {
			log.Printf("Error querying database for cart_id %d: %v", cartID, err)
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}
		log.Printf("Cart ID: %d, Quantity: %d", cartID, qty)
		weight += (qty * 1000)
	}

	log.Printf("Total weight calculated: %d grams", weight)
	weightGrams := weight

	deliveryReq := service.DeliveryCostReq{
		Origin:      req.OriginCityId,
		Destination: req.DestinationCityId,
		Weight:      fmt.Sprintf("%d", weightGrams),
		Courier:     req.CourierName,
	}
	deliveryResp, err := service.CallDeliveryAPI(deliveryReq)
	if err != nil {
		log.Printf("Error calling delivery API: %v\n", err)
		return nil, err
	}

	var index int = -1
	for i, service := range deliveryResp.Rajaongkir.Results[0].Costs {
		if service.Service == req.CourierServiceName {
			index = i
			break
		}
	}

	if index == -1 {
		log.Printf("Service '%s' not found.\n", req.CourierServiceName)
		return nil, status.Errorf(codes.NotFound, "Service '%s' not found", req.CourierServiceName)
	}

	deliveryFee := deliveryResp.Rajaongkir.Results[0].Costs[index].Cost[0].Value
	log.Printf("Calculated delivery fee: %d\n", deliveryFee)

	// Calculate total price
	fee := deliveryFee + int(req.OtherFee)
	totalPrice := price + fee - discount
	log.Printf("Calculated totalPrice: %d\n", totalPrice)

	XenditPayload, err := service.CallXenditInvoiceAPI(totalPrice)
	if err != nil {
		log.Printf("Error calling Xendit Invoice API: %v\n", err)
		return nil, err
	}

	// Insert user into database
	query = `INSERT INTO orders (user_id, voucher_id, status, price, fee, discount, total_price, metadata) VALUES (?, ?, ?, ?, ?, ?, ?,?)`
	result, err := tx.Exec(query, req.UserId, voucherID, "open", price, fee, discount, totalPrice, req.Metadata)
	if err != nil {
		log.Printf("Error inserting into orders table: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error inserting order: %v", err)
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error fetching order_id: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error fetching order_id: %v", err)
	}
	log.Printf("Order created with ID: %d\n", orderID)

	// Insert into payments
	query = `INSERT INTO payments (order_id, payment_external_id, amount, status, metadata) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.Exec(query, orderID, XenditPayload.Req.ExternalId, totalPrice, "open", req.Metadata)
	if err != nil {
		log.Printf("Error inserting into payments table: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error inserting payment: %v", err)
	}

	// Insert into deliveries
	query = `INSERT INTO deliveries (order_id, courier_name, courier_service, weight_grams, origin_city_id, destination_city_id, delivery_fee, status, metadata) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.Exec(query, orderID, req.CourierName, req.CourierServiceName, weightGrams, req.OriginCityId, req.DestinationCityId, deliveryFee, "open", req.Metadata)
	if err != nil {
		log.Printf("Error inserting into deliveries table: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error inserting delivery: %v", err)
	}

	// Insert into order details
	for _, shoe := range shoeList {
		query = `INSERT INTO order_details (order_id, shoe_id, quantity) VALUES (?, ?, ?)`
		_, err = tx.Exec(query, orderID, shoe.ID, shoe.Qty)
		if err != nil {
			log.Printf("Error inserting into order_details table: %v\n", err)
			return nil, status.Errorf(codes.Internal, "error inserting order details: %v", err)
		}
		log.Printf("Inserted order detail for Shoe ID: %d, Quantity: %d\n", shoe.ID, shoe.Qty)
	}

	// Set voucher expired
	if voucherID != "NOVOUCHER" {
		query = `UPDATE vouchers SET used = true WHERE voucher_id = ?`
		_, err = tx.Exec(query, voucherID)
		if err != nil {
			log.Printf("Error updating vouchers table: %v\n", err)
			return nil, status.Errorf(codes.Internal, "error updating voucher: %v", err)
		}
	}

	// Set cart empty
	for _, cartId := range cartIds {
		query = `DELETE FROM carts WHERE cart_id = ?`
		_, err = tx.Exec(query, cartId)
		if err != nil {
			log.Printf("Error deleting from carts table: %v\n", err)
			return nil, status.Errorf(codes.Internal, "error deleting cart: %v", err)
		}
	}

	// Prepare response
	response := &pb.AddOrderResponse{
		Message:     "Success",
		InvoiceUrl:  XenditPayload.Resp.InvoiceUrl,
		ExpiredDate: XenditPayload.Resp.ExpiryDate,
		TotalPrice:  int32(XenditPayload.Resp.Amount),
	}

	log.Println("Order successfully created")
	return response, nil
}

// https://developers.xendit.co/api-reference/#invoice-callback

func (h *OrderHandler) CallbackNotification(ctx context.Context, req *pb.CallbackNotificationRequest) (*empty.Empty, error) {
	return nil, nil
}
