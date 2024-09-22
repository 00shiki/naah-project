package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"order-service/config"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaymentReq struct {
}

type PaymentResp struct {
}

// # CallPaymentAPI
// CallPaymentAPI handles the HTTP request to the payment API and returns the response
func CallPaymentAPI(payload DeliveryCostReq) (*DeliveryCostResp, error) {
	// Convert payload to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return nil, status.Errorf(codes.Internal, "Error marshalling JSON: %v", err)
	}

	// Log the JSON request body
	log.Printf("Request body to 3rd party API: %s", string(body))

	// Create a new HTTP request
	reqBody := bytes.NewBuffer(body)
	httpReq, err := http.NewRequest(http.MethodPost, config.DELIVERY_COST_URL, reqBody)
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return nil, status.Errorf(codes.Internal, "Error creating HTTP request: %v", err)
	}

	// Set headers
	httpReq.Header.Set("key", config.DELIVERY_API_KEY)
	httpReq.Header.Set("Content-Type", "application/json")

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send the request
	log.Println("Sending HTTP request to delivery API")
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, status.Errorf(codes.Internal, "Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("API request failed with status: %s\nResponse body: %s", resp.Status, string(bodyBytes))
		return nil, status.Errorf(codes.Internal, "API request failed with status: %s", resp.Status)
	}

	// Parse the response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, status.Errorf(codes.Internal, "Error reading response body: %v", err)
	}

	// Log the response body
	log.Printf("Response body from 3rd party API: %s", string(bodyBytes))

	// Unmarshal the JSON response
	var apiResponse DeliveryCostResp // Define apiResponseStruct based on the expected response
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		log.Println("Error decoding response:", err)
		return nil, status.Errorf(codes.Internal, "Error decoding response: %v", err)
	}

	return &apiResponse, nil
}
