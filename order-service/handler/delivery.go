package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"order-service/config"
	"order-service/pb"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryHandler struct {
	db *sql.DB
}

func NewDeliveryHandler(db *sql.DB) *DeliveryHandler {
	return &DeliveryHandler{db}
}

type DeliveryCostReq struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      string `json:"weight"`
	Courier     string `json:"courier"`
}

func (h *DeliveryHandler) DeliveryCost(ctx context.Context, req *pb.DeliveryCostRequest) (*pb.DeliveryCostResponse, error) {
	log.Println("DeliveryCost method started")

	// ambil cart_id
	var weight int32
	var qty int32
	for i := 0; i < len(req.CartId); i++ {
		query := "SELECT quantity FROM cart WHERE cart_id = ?"
		log.Printf("Running query: %s with cart_id: %d", query, req.CartId[i])

		err := h.db.QueryRow(query, req.CartId[i]).Scan(&qty)
		if err != nil {
			log.Printf("Error querying database for cart_id %d: %v", req.CartId[i], err)
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}
		log.Printf("Cart ID: %d, Quantity: %d", req.CartId[i], qty)
		weight += (qty * 1000)
	}

	log.Printf("Total weight calculated: %d grams", weight)

	// Prepare the payload
	payload := DeliveryCostReq{
		Origin:      req.OriginCityId,
		Destination: req.DestinationCityId,
		Weight:      strconv.Itoa(int(weight)),
		Courier:     req.Courier,
	}

	log.Printf("Payload prepared: %+v", payload)

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

	// Parse and log the response body for debugging
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, status.Errorf(codes.Internal, "Error reading response body: %v", err)
	}
	log.Printf("Response body from 3rd party API: %s", string(bodyBytes))

	// Check for successful response
	log.Printf("Received HTTP response with status: %s", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status: %s\n", resp.Status)
		return nil, status.Errorf(codes.Internal, "API request failed with status: %s", resp.Status)
	}

	// Parse the response
	log.Println("Parsing API response")
	var apiResponse struct {
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

	// Unmarshal the JSON response
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		log.Println("Error decoding response:", err)
		return nil, status.Errorf(codes.Internal, "Error decoding response: %v", err)
	}

	log.Printf("Parsed API response: %+v", apiResponse)

	origin := pb.DeliveryItem{
		CityId:       apiResponse.Rajaongkir.OriginDetails.CityID,
		CityName:     apiResponse.Rajaongkir.OriginDetails.CityName,
		PostalCode:   apiResponse.Rajaongkir.OriginDetails.PostalCode,
		ProvinceName: apiResponse.Rajaongkir.OriginDetails.Province,
		ProvinceId:   apiResponse.Rajaongkir.OriginDetails.ProvinceID,
		Type:         apiResponse.Rajaongkir.OriginDetails.Type,
	}
	destination := pb.DeliveryItem{
		CityId:       apiResponse.Rajaongkir.DestinationDetails.CityID,
		CityName:     apiResponse.Rajaongkir.DestinationDetails.CityName,
		PostalCode:   apiResponse.Rajaongkir.DestinationDetails.PostalCode,
		ProvinceName: apiResponse.Rajaongkir.DestinationDetails.Province,
		ProvinceId:   apiResponse.Rajaongkir.DestinationDetails.ProvinceID,
		Type:         apiResponse.Rajaongkir.DestinationDetails.Type,
	}
	items := []*pb.ServiceItem{}

	for i := range apiResponse.Rajaongkir.Results[0].Costs {
		item := pb.ServiceItem{
			ServiceName: apiResponse.Rajaongkir.Results[0].Costs[i].Service,
			Description: apiResponse.Rajaongkir.Results[0].Costs[i].Description,
			Cost:        int32(apiResponse.Rajaongkir.Results[0].Costs[i].Cost[0].Value),
			Etd:         apiResponse.Rajaongkir.Results[0].Costs[i].Cost[0].Etd,
		}
		items = append(items, &item)
	}

	// Construct the response for gRPC
	log.Println("Constructing gRPC response")
	response := &pb.DeliveryCostResponse{
		Origin:      &origin,
		Destination: &destination,
		Service:     items,
	}

	log.Println("DeliveryCost method finished")
	return response, nil
}

func (h *DeliveryHandler) GetProvince(ctx context.Context, req *pb.Empty) (*pb.GetProvinceResponse, error) {
	log.Println("GetProvince method started")

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
	var apiResponse struct {
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

	// Decode the response body
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		log.Println("Error decoding response:", err)
		return nil, status.Errorf(codes.Internal, "Error decoding response: %v", err)
	}

	// Log the parsed API response for debugging
	log.Printf("Parsed API response: %+v", apiResponse)

	// Convert the response to the gRPC response format
	provinceList := []*pb.Province{}
	for _, result := range apiResponse.Rajaongkir.Results {
		provinceList = append(provinceList, &pb.Province{
			ProvinceId:   result.ProvinceID,
			ProvinceName: result.Province,
		})
	}

	// Return the gRPC response
	response := &pb.GetProvinceResponse{
		Provinces: provinceList,
	}

	log.Println("GetProvince method finished")
	return response, nil
}

func (h *DeliveryHandler) GetCity(ctx context.Context, req *pb.GetCityRequest) (*pb.GetCityResponse, error) {
	log.Println("GetCity method started")

	// Construct the URL with the province query parameter
	url := fmt.Sprintf("%s?province=%s", config.DELIVERY_CITY_URL, req.ProvinceId)

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
	var apiResponse struct {
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

	// Decode the response body
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		log.Println("Error decoding response:", err)
		return nil, status.Errorf(codes.Internal, "Error decoding response: %v", err)
	}

	// Log the parsed API response for debugging
	log.Printf("Parsed API response: %+v", apiResponse)

	// Convert the response to the gRPC response format
	cityList := []*pb.City{}
	for _, result := range apiResponse.Rajaongkir.Results {
		cityList = append(cityList, &pb.City{
			CityId:       result.CityID,
			CityName:     result.CityName,
			ProvinceId:   result.ProvinceID,
			ProvinceName: result.Province,
			Type:         result.Type,
			PostalCode:   result.PostalCode,
		})
	}

	// Return the gRPC response
	response := &pb.GetCityResponse{
		Cities: cityList,
	}

	log.Println("GetCity method finished")
	return response, nil
}
