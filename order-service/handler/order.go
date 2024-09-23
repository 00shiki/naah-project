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

	// Dummy values for now
	// userID := 1
	// voucherID := "NOVOUCHER"
	// price := 10000
	// otherFee := 0
	// fee := deliveryFee + otherFee
	// discount := 0
	// totalPrice := price + fee - discount

	/*
		 TODO
		- Buat method Callback Payment (tambah email ke pelanggan)
		- Buat method check order saja
	*/
	cartIds := utils.RemoveDuplicates(req.CartIds)

	var shoeId, shoeQty, shoePrice int
	var shoeList []shoes

	for _, cartId := range cartIds {
		query := "SELECT shoe_id, quantity FROM carts WHERE cart_id = ?"
		err := h.db.QueryRow(query, cartId).Scan(&shoeId, &shoeQty)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}

		query = `SELECT sm.price FROM shoe_models sm JOIN shoe_details sd ON sm.model_id = sd.model_id WHERE sd.shoe_id = ?`
		err = h.db.QueryRow(query, shoeId).Scan(&shoePrice)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}

		shoeList = append(shoeList, shoes{ID: shoeId, Price: shoePrice, Qty: shoeQty})
	}

	var price int
	for _, shoe := range shoeList {
		tempPrice := shoe.Price * shoe.Qty
		price += tempPrice
	}

	// TODO - Cek Harga sepatu grpc product
	/*
		TODO:
		create variable temp: total_price
		2. iterasi:
				. create variable temp: shoe_qty, shoe_price
				1. get cart_id
				- cek apakah user_id sesuai dengan user_id di cart
				- ambil shoes_id dan quantity jika benar
				2. get shoes detail by shoes_id -- grpc product
				- ambil shoe_price

				- add total_price += shoe_proce * shoe_qty
			 end iterasi
	*/
	//! CHECK VOUCHER !!!
	var voucherID string
	if req.VoucherId == "" {
		voucherID = "NOVOUCHER"
	} else {
		voucherID = req.VoucherId
	}
	// Variables to store data from the query
	var discountPercent float64
	var validUntilStr string // Store valid_until as a string initially
	var used bool

	// Execute the query
	query := "SELECT discount, valid_until, used FROM vouchers WHERE voucher_id = ?"
	log.Printf("Executing query: %s\n", query)
	err := h.db.QueryRow(query, req.VoucherId).Scan(&discountPercent, &validUntilStr, &used)
	if err != nil {
		if err == sql.ErrNoRows {
			// No voucher found for the given voucher ID
			log.Printf("No voucher found for VoucherId=%s\n", req.VoucherId)
			return nil, status.Errorf(codes.NotFound, "Voucher not found")
		} else {
			// Handle unexpected database errors and return a gRPC Internal error
			log.Printf("Error querying database for VoucherId=%s: %v\n", req.VoucherId, err)
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}
	}

	// Parse the validUntil string into time.Time
	layout := "2006-01-02" // Adjust this layout to match the format in your database
	validUntil, err := time.Parse(layout, validUntilStr)
	if err != nil {
		log.Printf("Error parsing validUntil: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Error parsing validUntil: %v", err)
	}

	// Check if the voucher is valid
	if validUntil.Before(time.Now()) {
		return nil, status.Errorf(codes.InvalidArgument, "Voucher has expired")
	} else if used {
		return nil, status.Errorf(codes.InvalidArgument, "Voucher has already used")
	}

	// Assuming `price` is an int representing the price in cents
	// Calculate the discount as an integer
	discount := int(float64(price) * (discountPercent / 100)) // Calculate discount in cents
	log.Printf("Applying discount: %d on price: %d\n", discount, price)

	//! CHECK COURIER !!!
	// Calculate total weight
	var weight int32
	for _, cartID := range req.CartIds {
		query := "SELECT quantity FROM carts WHERE cart_id = ?"
		log.Printf("Running query: %s with cart_id: %d", query, cartID)

		var qty int32
		err := h.db.QueryRow(query, cartID).Scan(&qty)
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
		Weight:      fmt.Sprintf("%s", weightGrams),
		Courier:     req.CourierName,
	}
	deliveryResp, err := service.CallDeliveryAPI(deliveryReq)
	if err != nil {
		return nil, err
	}

	// Find the index of the service with ServiceName
	var index int = -1
	for i, service := range deliveryResp.Rajaongkir.Results[0].Costs {
		if service.Service == req.CourierServiceName {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Service '%s' not found.\n", req.CourierServiceName)
		return nil, status.Errorf(codes.NotFound, "Service '%s' not found", req.CourierServiceName)
	}

	deliveryFee := deliveryResp.Rajaongkir.Results[0].Costs[index].Cost[0].Value

	//! TOTAL
	fee := deliveryFee + int(req.OtherFee)
	totalPrice := price + fee - discount

	//! MAKE ORDER !!!
	// Log some important variables
	log.Printf("Calculated totalPrice: %d\n", totalPrice)

	XenditPayload, err := service.CallXenditInvoiceAPI(totalPrice)
	if err != nil {
		return nil, err
	}

	// Insert user into database
	query = `INSERT INTO orders (user_id, voucher_id, status, price, fee, discount, total_price, metadata) VALUES (?, ?, ?, ?, ?, ?, ?,?)`
	result, err := h.db.Exec(query, req.UserId, voucherID, "open", price, fee, discount, totalPrice, req.Metadata)
	if err != nil {
		log.Printf("Error inserting into orders table: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error inserting order: %v", err)
	}

	// Retrieve the last inserted order_id
	orderID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error fetching order_id: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error fetching order_id: %v", err)
	}

	log.Printf("Order created with ID: %d\n", orderID)

	// Insert into payments
	query = `INSERT INTO payments (order_id, payment_external_id, amount, status, metadata) VALUES (?, ?, ?, ?, ?)`
	_, err = h.db.Exec(query, orderID, XenditPayload.Req.ExternalId, totalPrice, "open", req.Metadata)
	if err != nil {
		log.Printf("Error inserting into payments table: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error inserting payment: %v", err)
	}

	// Insert into deliveries
	query = `INSERT INTO deliveries (order_id, courier_name, courier_service, weight_grams, origin_city_id, destination_city_id, delivery_fee, status, metadata) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = h.db.Exec(query, orderID, req.CourierName, req.CourierServiceName, weightGrams, req.OriginCityId, req.DestinationCityId, deliveryFee, "open", req.Metadata)
	if err != nil {
		log.Printf("Error inserting into deliveries table: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error inserting delivery: %v", err)
	}

	// TODO Insert Into order details

	// Insert into order details
	for index, shoeId := range shoeIds {
		query = `INSERT INTO order_details (order_id, shoe_id, quantity) VALUES (?, ?, ?)`
		_, err = h.db.Exec(query, orderID, shoeId, qtys[index])
		if err != nil {
			log.Printf("Error inserting into deliveries table: %v\n", err)
			return nil, status.Errorf(codes.Internal, "error inserting delivery: %v", err)
		}
	}
	// TODO disable voucher after used

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
