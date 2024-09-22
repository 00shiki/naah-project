package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io"
	"math"
	"net/http"
	"order-service/config"
	"order-service/pb"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderHandler struct {
	db *sql.DB
}

type XenditInvoiceReq struct {
	ExternalId string
	Amount     int
}

type XenditInvoiceResp struct {
	Status     string
	Amount     int
	InvoiceUrl string
	ExpiryDate string
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{db}
}

func (h *OrderHandler) AddOrder(ctx context.Context, req *empty.Empty) (*pb.AddOrderResponse, error) {
	/*
				TODO:
				1. get request
				- user_id
				- list of cart_id
				- voucher_id
				- delivery
				"origin_city_id"
		    "destination_city_id"
		    "courier"
				"courier_service_name"

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

				3. Create Order
				- Orders table
				- order_details table
				  iterasi:
					input
					- order_id
					- shoe_id & quantity

				4. Cek Delivery
				- Cek harga delivery based on courier
				- add total_price

				5. Create Payment
				-






	*/
	//Dummy
	userID := 1
	voucherID := "NOVOUCHER"
	price := 10000
	deliveryFee := 20000
	otherFee := 0
	fee := deliveryFee + otherFee
	discount := 0
	totalPrice := price + fee - discount
	courierName := "jne"
	courierServiceName := "REG"
	weightGrams := 2000
	originCityId := "1"
	destinationCityId := "1"
	metadata := "nothing"

	// Generate a new UUID and set up the invoice request body
	xenditInvoiceReq := XenditInvoiceReq{
		ExternalId: uuid.New().String(),
		Amount:     totalPrice,
	}

	// Marshal request body into JSON
	reqBody, err := json.Marshal(xenditInvoiceReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Marshall request Body: %v", err)
	}

	invoiceURL := config.XENDIT_INVOICE_URL
	xenditAPIKey := config.XENDIT_SECRET_KEY
	client := &http.Client{}
	httpReq, err := http.NewRequest("POST", invoiceURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create request: %v", err)
	}

	httpReq.Header.Set("Authorization", "Basic "+basicAuth(xenditAPIKey))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read and Unmarshal the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to read response: %v", err)
	}

	var responseXendit map[string]interface{}
	if err := json.Unmarshal(respBody, &responseXendit); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to parse response: %v", err)
	}

	// Check for errors in the response
	if errorCode, ok := responseXendit["error_code"].(string); ok {
		errorMessage := "An error occurred"
		if msg, ok := responseXendit["message"].(string); ok {
			errorMessage = msg
		}

		// Return error response
		return nil, status.Errorf(codes.Internal, "error from partner response: %v %v", errorCode, errorMessage)
	}

	// Extract only the required fields if no errors
	xenditInvoiceResp := XenditInvoiceResp{
		Status:     responseXendit["status"].(string),
		Amount:     int(responseXendit["amount"].(float64)), // Safely assert as float64 and convert to int
		InvoiceUrl: responseXendit["invoice_url"].(string),
		ExpiryDate: responseXendit["expiry_date"].(string),
	}

	if xenditInvoiceResp.Status != "PENDING" {
		return nil, status.Errorf(codes.Internal, "error from partner response")
	} else if xenditInvoiceResp.Amount != totalPrice {
		return nil, status.Errorf(codes.Internal, "error total price from partner response")
	}

	// Insert user into database
	query := `INSERT INTO orders (user_id, voucher_id, status, price, fee, discount, total_price) VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := h.db.Exec(query, userID, voucherID, "open", price, fee, discount, totalPrice)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error from partner response: %v", err)
	}

	// Retrieve the last inserted order_id
	orderID, err := result.LastInsertId()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error fetching order_id: %v", err)
	}

	query = `INSERT INTO payments (order_id, payment_external_id, amount, status) VALUES (?, ?, ?, ?)`
	_, err = h.db.Exec(query,
		orderID, xenditInvoiceReq.ExternalId, totalPrice, "open")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error from partner response: %v", err)
	}

	query = `INSERT INTO deliveries (order_id, courier_name, courier_service, weight_grams, origin_city_id, destination_city_id, delivery_fee, status, metadata) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = h.db.Exec(query,
		orderID, courierName, courierServiceName, weightGrams, originCityId, destinationCityId, deliveryFee, "open", metadata)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error from partner response: %v", err)
	}

	response := &pb.AddOrderResponse{
		Message:     "Success",
		InvoiceUrl:  xenditInvoiceResp.InvoiceUrl,
		ExpiredDate: xenditInvoiceResp.ExpiryDate,
		TotalPrice:  int32(math.Round(responseXendit["amount"].(float64))),
	}

	return response, nil
}

func basicAuth(apiKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))
}
