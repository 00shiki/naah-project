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
	db  *sql.DB
	rmq *utils.RabbitMQClient
}

func NewOrderHandler(db *sql.DB, rmq *utils.RabbitMQClient) *OrderHandler {
	return &OrderHandler{db, rmq}
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

	// Calculate total price for shoes
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
	if req.OtherFee < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid otherFee amount")
	}
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

	for _, shoe := range shoeList {
		// Get the current stock
		var currentStock int
		query := `SELECT stock FROM shoe_details WHERE shoe_id = ?`
		err := tx.QueryRow(query, shoe.ID).Scan(&currentStock)
		if err != nil {
			log.Printf("Error fetching stock for shoe_id %d: %v\n", shoe.ID, err)
			return nil, status.Errorf(codes.Internal, "error fetching stock for shoe ID %d: %v", shoe.ID, err)
		}

		// Check if reducing the stock will make it go below 0
		if currentStock < shoe.Qty {
			log.Printf("Insufficient stock for Shoe ID: %d. Available stock: %d, Requested quantity: %d\n", shoe.ID, currentStock, shoe.Qty)
			return nil, status.Errorf(codes.FailedPrecondition, "insufficient stock for Shoe ID %d", shoe.ID)
		}

		// Insert into order_details
		query = `INSERT INTO order_details (order_id, shoe_id, quantity) VALUES (?, ?, ?)`
		_, err = tx.Exec(query, orderID, shoe.ID, shoe.Qty)
		if err != nil {
			log.Printf("Error inserting into order_details table: %v\n", err)
			return nil, status.Errorf(codes.Internal, "error inserting order details: %v", err)
		}
		log.Printf("Inserted order detail for Shoe ID: %d, Quantity: %d\n", shoe.ID, shoe.Qty)

		// Update stock
		query = `UPDATE shoe_details SET stock = stock - ? WHERE shoe_id = ?`
		_, err = tx.Exec(query, shoe.Qty, shoe.ID)
		if err != nil {
			log.Printf("Error updating shoe_details table: %v\n", err)
			return nil, status.Errorf(codes.Internal, "error updating stock: %v", err)
		}
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
// type OrderReceiptEmail struct {
// 	UserName    string        `json:"user_name"`
// 	UserEmail   string        `json:"user_email"`
// 	OrderId     int64         `json:"order_id"`
// 	TotalPrice  int32         `json:"total_price"`
// 	OrderDetail []OrderDetail `json:"order_detail"`
// }

func (h *OrderHandler) CallbackNotification(ctx context.Context, req *pb.CallbackNotificationRequest) (*empty.Empty, error) {
	log.Printf("Starting CallbackNotification for payment_external_id: %s, paid_amount: %d, status: %s", req.OrderIdExt, req.PaidAmount, req.Status)

	// Start a transaction
	tx, err := h.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Printf("Error starting transaction: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error starting transaction: %v", err)
	}

	// Defer a rollback in case something goes wrong
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback() // If panic, rollback
			log.Printf("Transaction panicked, rolling back: %v\n", p)
			panic(p) // Rethrow panic after rollback
		} else if err != nil {
			log.Printf("Transaction rolled back due to error: %v\n", err)
			tx.Rollback() // If error exists, rollback
		} else {
			err = tx.Commit() // If all went well, commit
			if err != nil {
				log.Printf("Error committing transaction: %v\n", err)
			} else {
				log.Println("Transaction committed successfully")
			}
		}
	}()

	// Check if the paid amount matches the payment's amount
	var paymentAmount, orderID int32
	query := `SELECT amount, order_id FROM payments WHERE payment_external_id = ?`
	log.Printf("Executing query: %s", query)
	err = tx.QueryRowContext(ctx, query, req.OrderIdExt).Scan(&paymentAmount, &orderID)
	if err != nil {
		log.Printf("Error fetching payment details for payment_external_id %s: %v\n", req.OrderIdExt, err)
		return nil, status.Errorf(codes.Internal, "error fetching payment details: %v", err)
	}

	log.Printf("Fetched payment details: payment_amount=%d, order_id=%d", paymentAmount, orderID)

	if req.PaidAmount != paymentAmount {
		log.Printf("Paid amount mismatch: expected=%d, received=%d", paymentAmount, req.PaidAmount)
		return nil, status.Errorf(codes.InvalidArgument, "paid amount does not match the expected amount")
	}

	// Update the payment status
	query = `UPDATE payments SET status = ?, updated_at = NOW() WHERE payment_external_id = ?`
	log.Printf("Executing query to update payment status: %s", query)
	_, err = tx.ExecContext(ctx, query, req.Status, req.OrderIdExt)
	if err != nil {
		log.Printf("Error updating payment status for payment_external_id %s: %v\n", req.OrderIdExt, err)
		return nil, status.Errorf(codes.Internal, "error updating payment status: %v", err)
	}
	log.Println("Payment status updated successfully")

	// Update the order status
	query = `UPDATE orders SET status = ?, updated_at = NOW() WHERE order_id = ?`
	log.Printf("Executing query to update order status: %s", query)
	_, err = tx.ExecContext(ctx, query, req.Status, orderID)
	if err != nil {
		log.Printf("Error updating order status for order_id %d: %v\n", orderID, err)
		return nil, status.Errorf(codes.Internal, "error updating order status: %v", err)
	}
	log.Println("Order status updated successfully")

	// Fetch order receipt details using the provided SQL query
	var userNameFirst, userNameLast, userEmail, shoeModelName string
	var shoeSize int64
	var quantity, totalPrice int32
	orderDetails := []utils.OrderDetail{}

	log.Printf("Executing query to fetch order receipt details for payment_external_id %s", req.OrderIdExt)
	rows, err := tx.QueryContext(ctx, `
		SELECT u.first_name, u.last_name, u.email, p.order_id, p.amount, sm.name AS shoe_model_name, sd.size AS shoe_size, od.quantity
		FROM payments p
		JOIN orders o ON p.order_id = o.order_id
		JOIN users u ON o.user_id = u.user_id
		JOIN order_details od ON o.order_id = od.order_id
		JOIN shoe_details sd ON od.shoe_id = sd.shoe_id
		JOIN shoe_models sm ON sd.model_id = sm.model_id
		WHERE p.payment_external_id = ?`, req.OrderIdExt)

	if err != nil {
		log.Printf("Error fetching order receipt details for payment_external_id %s: %v\n", req.OrderIdExt, err)
		return nil, status.Errorf(codes.Internal, "error fetching order receipt details: %v", err)
	}
	defer rows.Close()

	log.Println("Processing order receipt details")
	// Loop through results and gather order details
	for rows.Next() {
		err := rows.Scan(&userNameFirst, &userNameLast, &userEmail, &orderID, &totalPrice, &shoeModelName, &shoeSize, &quantity)
		if err != nil {
			log.Printf("Error scanning order details: %v\n", err)
			return nil, status.Errorf(codes.Internal, "error scanning order details: %v", err)
		}
		orderDetails = append(orderDetails, utils.OrderDetail{
			ShoeName: shoeModelName,
			ShoeSize: shoeSize,
			Qty:      quantity,
		})
		log.Printf("Scanned order detail: %s, size=%d, quantity=%d", shoeModelName, shoeSize, quantity)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error processing order details: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error processing order details: %v", err)
	}

	// Create the order receipt email payload
	orderReceipt := utils.OrderReceiptEmail{
		UserName:    fmt.Sprintf("%s %s", userNameFirst, userNameLast),
		UserEmail:   userEmail,
		OrderId:     int64(orderID),
		TotalPrice:  totalPrice,
		OrderDetail: orderDetails,
	}
	log.Printf("Order receipt email created for user: %s (%s)", orderReceipt.UserName, userEmail)

	// Send the receipt email using RabbitMQ
	err = utils.SendReceiptEmail(orderReceipt, h.rmq)
	if err != nil {
		log.Printf("Error sending order receipt email for order_id %d: %v\n", orderID, err)
		return nil, status.Errorf(codes.Internal, "error sending order receipt email: %v", err)
	}
	log.Println("Order receipt email sent successfully")

	return &empty.Empty{}, nil
}

func (h *OrderHandler) GetOrderList(ctx context.Context, req *pb.GetOrderListRequest) (*pb.GetOrderListResponse, error) {
	userId := req.UserId
	log.Printf("Received GetOrderList request for user ID: %d", userId)

	query := `
        SELECT 
            o.order_id,
            o.status,
            o.fee,
            o.discount,
            o.total_price,
            o.voucher_id,
            sd.size AS shoe_size,
            sm.name AS shoe_model_name,
            sm.price AS shoe_model_price,
            od.quantity AS shoe_qty
        FROM 
            orders o
        LEFT JOIN 
            order_details od ON o.order_id = od.order_id
        LEFT JOIN 
            shoe_details sd ON od.shoe_id = sd.shoe_id
        LEFT JOIN
            shoe_models sm ON sd.model_id = sm.model_id
        WHERE 
            o.user_id = ?
        ORDER BY 
            o.order_id;`

	log.Println("Executing SQL query...")
	rows, err := h.db.Query(query, userId)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int32]*pb.Order)

	for rows.Next() {
		var order pb.Order
		var shoe pb.Shoe

		// Scan the values from the query
		err = rows.Scan(
			&order.OrderId,
			&order.Status,
			&order.Fee,
			&order.Discount,
			&order.TotalPrice,
			&order.VoucherId,
			&shoe.Size,
			&shoe.Name,
			&shoe.Price,
			&shoe.Qty,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		// Check if the order already exists in the map
		existingOrder, found := ordersMap[(order.OrderId)]
		if !found {
			// If not found, create a new order entry
			existingOrder = &order
			existingOrder.Shoes = []*pb.Shoe{} // Initialize the shoes list
			ordersMap[order.OrderId] = existingOrder
		}

		// Append the shoe to the existing order's shoe list
		existingOrder.Shoes = append(existingOrder.Shoes, &shoe)

		// Log the processed order
		log.Printf("Processed order: %+v", existingOrder)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error encountered during row iteration: %v", err)
		return nil, err
	}

	// Convert ordersMap to a slice
	var orders []*pb.Order
	for _, order := range ordersMap {
		orders = append(orders, order)
	}

	log.Printf("Returning %d orders for user ID: %d", len(orders), userId)
	return &pb.GetOrderListResponse{Orders: orders}, nil
}
