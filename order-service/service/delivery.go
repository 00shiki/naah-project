package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"order-service/config"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryCostReq struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      string `json:"weight"`
	Courier     string `json:"courier"`
}

type DeliveryCostResp struct {
	Rajaongkir struct {
		Query struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Weight      int    `json:"weight"`
			Courier     string `json:"courier"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		OriginDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"origin_details"`
		DestinationDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"destination_details"`
		Results []struct {
			Code  string `json:"code"`
			Name  string `json:"name"`
			Costs []struct {
				Service     string `json:"service"`
				Description string `json:"description"`
				Cost        []struct {
					Value int    `json:"value"`
					Etd   string `json:"etd"`
					Note  string `json:"note"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

type provinceResponseStruct struct {
	Rajaongkir struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []struct {
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

type cityResponseStruct struct {
	Rajaongkir struct {
		Query struct {
			Province string `json:"province"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

// # CallDeliveryAPI
// CallDeliveryAPI handles the HTTP request to the delivery API and returns the response
func CallDeliveryAPI(payload DeliveryCostReq) (*DeliveryCostResp, error) {
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

// # CallGetProvinceAPI
// CallGetProvinceAPI makes the HTTP request to the delivery API
func CallGetProvinceAPI() (*provinceResponseStruct, error) {
	// Create a new HTTP request
	httpReq, err := http.NewRequest(http.MethodGet, config.DELIVERY_PROVINCE_URL, nil)
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return nil, status.Errorf(codes.Internal, "Error creating HTTP request: %v", err)
	}

	// Set headers
	httpReq.Header.Set("key", config.DELIVERY_API_KEY)
	httpReq.Header.Set("Content-Type", "application/json")

	// Log the request for debugging
	log.Printf("HTTP Request: %v", httpReq)

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for the request
	}

	// Send the request
	log.Println("Sending HTTP request to delivery API")
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, status.Errorf(codes.Internal, "Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Log the HTTP status code
	log.Printf("Received HTTP response with status: %s", resp.Status)

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		// Read the response body for debugging if status is not OK
		body, _ := io.ReadAll(resp.Body)
		log.Printf("API request failed with status: %s\nResponse body: %s", resp.Status, string(body))
		return nil, status.Errorf(codes.Internal, "API request failed with status: %s", resp.Status)
	}

	// Parse and log the response body for debugging
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, status.Errorf(codes.Internal, "Error reading response body: %v", err)
	}
	log.Printf("Response body from 3rd party API: %s", string(bodyBytes))

	// Parse the response body into the API response struct
	var apiResponse provinceResponseStruct

	// Decode the response body
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		log.Println("Error decoding response:", err)
		return nil, status.Errorf(codes.Internal, "Error decoding response: %v", err)
	}

	// Log the parsed API response for debugging
	log.Printf("Parsed API response: %+v", apiResponse)

	return &apiResponse, nil
}

// #CallGetCityAPI
// CallGetCityAPI makes the HTTP request to the delivery API for cities
func CallGetCityAPI(provinceID string) (*cityResponseStruct, error) {
	// Construct the URL with the province query parameter
	url := fmt.Sprintf("%s?province=%s", config.DELIVERY_CITY_URL, provinceID)

	// Create a new HTTP request
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return nil, status.Errorf(codes.Internal, "Error creating HTTP request: %v", err)
	}

	// Set headers
	httpReq.Header.Set("key", config.DELIVERY_API_KEY)
	httpReq.Header.Set("Content-Type", "application/json")

	// Log the request for debugging
	log.Printf("HTTP Request: %v", httpReq)

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for the request
	}

	// Send the request
	log.Println("Sending HTTP request to delivery API")
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, status.Errorf(codes.Internal, "Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Log the HTTP status code
	log.Printf("Received HTTP response with status: %s", resp.Status)

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		// Read the response body for debugging if status is not OK
		body, _ := io.ReadAll(resp.Body)
		log.Printf("API request failed with status: %s\nResponse body: %s", resp.Status, string(body))
		return nil, status.Errorf(codes.Internal, "API request failed with status: %s", resp.Status)
	}

	// Parse and log the response body for debugging
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, status.Errorf(codes.Internal, "Error reading response body: %v", err)
	}
	log.Printf("Response body from 3rd party API: %s", string(bodyBytes))

	// Parse the response body into the API response struct
	var apiResponse cityResponseStruct

	// Decode the response body
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		log.Println("Error decoding response:", err)
		return nil, status.Errorf(codes.Internal, "Error decoding response: %v", err)
	}

	// Log the parsed API response for debugging
	log.Printf("Parsed API response: %+v", apiResponse)

	return &apiResponse, nil
}
