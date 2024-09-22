package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"order-service/config"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type XenditPayload struct {
	Req  XenditInvoiceReq
	Resp XenditInvoiceResp
}
type XenditInvoiceReq struct {
	ExternalId string `json:"external_id"`
	Amount     int    `json:"amount"`
}

type XenditInvoiceResp struct {
	Status     string
	Amount     int
	InvoiceUrl string
	ExpiryDate string
}

func CallXenditInvoiceAPI(totalPrice int) (payload *XenditPayload, err error) {

	// Generate a new UUID and set up the invoice request body
	xenditInvoiceReq := XenditInvoiceReq{
		ExternalId: uuid.New().String(),
		Amount:     totalPrice,
	}

	// Marshal request body into JSON
	reqBody, err := json.Marshal(xenditInvoiceReq)
	if err != nil {
		log.Printf("Error marshaling request body: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Marshall request Body: %v", err)
	}

	invoiceURL := config.XENDIT_INVOICE_URL
	xenditAPIKey := config.XENDIT_SECRET_KEY
	client := &http.Client{}
	httpReq, err := http.NewRequest("POST", invoiceURL, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Failed to create request: %v", err)
	}

	// Set headers and log them
	httpReq.Header.Set("Authorization", "Basic "+basicAuth(xenditAPIKey))
	httpReq.Header.Set("Content-Type", "application/json")
	log.Printf("Sending request to Xendit with URL: %s\n", invoiceURL)

	resp, err := client.Do(httpReq)
	if err != nil {
		log.Printf("Error sending HTTP request: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read and Unmarshal the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Failed to read response: %v", err)
	}

	// Log the raw response
	log.Printf("Raw response from Xendit: %s\n", string(respBody))

	var responseXendit map[string]interface{}
	if err := json.Unmarshal(respBody, &responseXendit); err != nil {
		log.Printf("Error unmarshaling response: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Failed to parse response: %v", err)
	}

	// Check for errors in the response
	if errorCode, ok := responseXendit["error_code"].(string); ok {
		errorMessage := "An error occurred"
		if msg, ok := responseXendit["message"].(string); ok {
			errorMessage = msg
		}
		log.Printf("Error from Xendit response: %s - %s\n", errorCode, errorMessage)
		return nil, status.Errorf(codes.Internal, "error from partner response: %v %v", errorCode, errorMessage)
	}

	// Extract only the required fields if no errors
	xenditInvoiceResp := XenditInvoiceResp{
		Status:     responseXendit["status"].(string),
		Amount:     int(responseXendit["amount"].(float64)), // Safely assert as float64 and convert to int
		InvoiceUrl: responseXendit["invoice_url"].(string),
		ExpiryDate: responseXendit["expiry_date"].(string),
	}

	// Log the extracted response
	log.Printf("Xendit response: Status: %s, Amount: %d, Invoice URL: %s\n", xenditInvoiceResp.Status, xenditInvoiceResp.Amount, xenditInvoiceResp.InvoiceUrl)

	if xenditInvoiceResp.Status != "PENDING" {
		log.Println("Xendit response status is not PENDING")
		return nil, status.Errorf(codes.Internal, "error from partner response")
	} else if xenditInvoiceResp.Amount != totalPrice {
		log.Printf("Total price mismatch: Expected %d, Got %d\n", totalPrice, xenditInvoiceResp.Amount)
		return nil, status.Errorf(codes.Internal, "error total price from partner response")
	}

	return &XenditPayload{
		xenditInvoiceReq,
		xenditInvoiceResp,
	}, nil
}

func basicAuth(apiKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))
}
