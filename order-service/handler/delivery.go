package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"order-service/config"
	"order-service/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryHandler struct {
	db *sql.DB
}

// type OrderItem struct {
// }

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

	// ambil cart_id
	var weight int32
	var qty int32
	for i := 0; i < len(req.CartId); i++ {
		query := "SELECT quantity FROM cart WHERE cart_id = ?"
		err := h.db.QueryRow(query, req.CartId[i]).Scan(&qty)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error querying database: %v", err)
		}
		weight += (qty * 1000)
	}

	// Prepare the payload
	payload := DeliveryCostReq{
		Origin:      req.OriginCityId,
		Destination: req.DestinationCityId,
		Weight:      string(weight),
		Courier:     req.Courier,
	}

	// Convert payload to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return nil, status.Errorf(codes.Internal, "Error marshalling JSON: %v", err)
	}

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
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, status.Errorf(codes.Internal, "Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status: %s\n", resp.Status)
		return nil, status.Errorf(codes.Internal, "API request failed with status: %s", resp.Status)
	}

	// Parse the response
	var apiResponse struct {
		Rajaongkir struct {
			Query struct {
				Test string `json:"test"`
				// Origin      string `json:"origin"`
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

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Println("Error decoding response:", err)
		return nil, status.Errorf(codes.Internal, "Error decoding response: %v", err)
	}

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

	// Construct the response for gRPC
	response := &pb.DeliveryCostResponse{
		Origin:      &origin,
		Destination: &destination,
		Service:     items,
		// Populate this with the necessary fields based on your protobuf definition
		// You might want to map the apiResponse fields into your DeliveryCostResponse
	}

	return response, nil
}

func (h *DeliveryHandler) GetProvince(ctx context.Context, req *pb.Empty) (*pb.GetProvinceResponse, error) {
	return nil, nil
}
func (h *DeliveryHandler) GetCity(ctx context.Context, req *pb.GetCityRequest) (*pb.GetCityResponse, error) {
	return nil, nil
}
