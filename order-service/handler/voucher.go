package handler

import (
	"context"
	"database/sql"
	"log"
	"order-service/pb"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VoucherHandler struct {
	db *sql.DB
}

func NewVoucherHandler(db *sql.DB) *VoucherHandler {
	return &VoucherHandler{db}
}

// service VoucherService {
//     rpc AddVoucher(AddVoucherRequest) returns (AddVoucherResponse);
//     rpc GetVoucher(GetVoucherRequest) returns (GetVoucherResponse);
//     rpc GetVoucherList(google.protobuf.Empty) returns (GetVoucherListResponse);
// }

func (h *VoucherHandler) AddVoucher(ctx context.Context, req *pb.AddVoucherRequest) (*pb.AddVoucherResponse, error) {
	// Log the received request
	log.Printf("AddVoucher request received: VoucherId=%s, Discount=%f, ValidUntil=%s\n", req.VoucherId, req.Discount, req.ValidUntil)

	// Define the layout format that matches the incoming validUntil string
	layout := "2006-01-02" // Use the layout that corresponds to the format of your date string

	// Parse the validUntil string into time.Time
	validUntil, err := time.Parse(layout, req.ValidUntil)
	if err != nil {
		log.Printf("Error parsing validUntil: %v\n", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid validUntil format: %v", err)
	}

	// Insert voucher into the database
	query := `INSERT INTO vouchers (voucher_id, discount, valid_until, used) VALUES (?, ?, ?, ?)`
	log.Printf("Executing query: %s\n", query)
	_, err = h.db.Exec(query, req.VoucherId, req.Discount, validUntil, false)
	if err != nil {
		log.Printf("Error inserting into vouchers table: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error inserting voucher: %v", err)
	}

	// Log successful addition
	log.Printf("Voucher %s added successfully\n", req.VoucherId)

	response := pb.AddVoucherResponse{
		Message: "Voucher Added",
	}

	return &response, nil
}

func (h *VoucherHandler) GetVoucher(ctx context.Context, req *pb.GetVoucherRequest) (*pb.GetVoucherResponse, error) {
	log.Printf("GetVoucher request received: VoucherId=%s\n", req.VoucherId)

	// Variables to store data from the query
	var discount float64
	var validUntilStr string // Store valid_until as a string initially
	var used bool

	// Execute the query
	query := "SELECT discount, valid_until, used FROM vouchers WHERE voucher_id = ?"
	log.Printf("Executing query: %s\n", query)
	err := h.db.QueryRow(query, req.VoucherId).Scan(&discount, &validUntilStr, &used)
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

	// Prepare the response
	response := pb.GetVoucherResponse{
		VoucherId:  req.VoucherId,
		Discount:   discount,
		ValidUntil: validUntil.Format("2006-01-02"), // Example format: "YYYY-MM-DD"
		Used:       used,
	}

	log.Printf("Voucher found: VoucherId=%s, Discount=%f, ValidUntil=%s, Used=%t\n", response.VoucherId, response.Discount, response.ValidUntil, response.Used)
	return &response, nil
}

func (h *VoucherHandler) GetVoucherList(ctx context.Context, req *empty.Empty) (*pb.GetVoucherListResponse, error) {
	// Log the received request
	log.Println("GetVoucherList request received")

	// Query to fetch all vouchers
	query := "SELECT voucher_id, discount, valid_until, used FROM vouchers"
	log.Printf("Executing query: %s\n", query)
	rows, err := h.db.Query(query)
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
	}
	defer rows.Close()

	// List to store all the vouchers
	var vouchers []*pb.Voucher

	// Loop through each row
	for rows.Next() {
		var voucher pb.Voucher
		var validUntilStr string // Store validUntil as string initially

		// Scan the row data into the voucher
		err := rows.Scan(&voucher.VoucherId, &voucher.Discount, &validUntilStr, &voucher.Used)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, status.Errorf(codes.Internal, "Error reading voucher data: %v", err)
		}

		// Parse the validUntil string into time.Time
		layout := "2006-01-02" // Adjust this layout based on your date format in the database
		validUntil, err := time.Parse(layout, validUntilStr)
		if err != nil {
			log.Printf("Error parsing validUntil: %v\n", err)
			return nil, status.Errorf(codes.Internal, "Error parsing validUntil: %v", err)
		}

		// Convert validUntil to string format
		voucher.ValidUntil = validUntil.Format("2006-01-02") // Example format: "YYYY-MM-DD"
		log.Printf("Voucher found: VoucherId=%s, Discount=%f, ValidUntil=%s, Used=%t\n", voucher.VoucherId, voucher.Discount, voucher.ValidUntil, voucher.Used)

		// Append the voucher to the list
		vouchers = append(vouchers, &voucher)
	}

	// Check for errors that occurred during row iteration
	if err = rows.Err(); err != nil {
		log.Println("Row iteration error:", err)
		return nil, status.Errorf(codes.Internal, "Row iteration error: %v", err)
	}

	// Create the response and return the list of vouchers
	response := pb.GetVoucherListResponse{
		Vouchers: vouchers,
	}

	log.Printf("Total vouchers found: %d\n", len(vouchers))
	return &response, nil
}
