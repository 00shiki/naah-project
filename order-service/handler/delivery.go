package handler

import (
	"context"
	"database/sql"
	"log"
	"order-service/pb"
	"order-service/service"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryHandler struct {
	db *sql.DB
}

func NewDeliveryHandler(db *sql.DB) *DeliveryHandler {
	return &DeliveryHandler{db}
}

// DeliveryCost method
func (h *DeliveryHandler) DeliveryCost(ctx context.Context, req *pb.DeliveryCostRequest) (*pb.DeliveryCostResponse, error) {
	log.Println("DeliveryCost method started")

	// Calculate total weight
	var weight int32
	for _, cartID := range req.CartId {
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

	// Prepare the payload
	payload := service.DeliveryCostReq{
		Origin:      req.OriginCityId,
		Destination: req.DestinationCityId,
		Weight:      strconv.Itoa(int(weight)),
		Courier:     req.Courier,
	}

	// Call the API
	apiResponse, err := service.CallDeliveryAPI(payload)
	if err != nil {
		return nil, err
	}

	// Construct gRPC response
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
	items := make([]*pb.ServiceItem, len(apiResponse.Rajaongkir.Results[0].Costs))

	for i, cost := range apiResponse.Rajaongkir.Results[0].Costs {
		items[i] = &pb.ServiceItem{
			ServiceName: cost.Service,
			Description: cost.Description,
			Cost:        int32(cost.Cost[0].Value),
			Etd:         cost.Cost[0].Etd,
		}
	}

	response := &pb.DeliveryCostResponse{
		Origin:      &origin,
		Destination: &destination,
		Service:     items,
	}

	log.Println("DeliveryCost method finished")
	return response, nil
}

func (h *DeliveryHandler) GetProvince(ctx context.Context, req *empty.Empty) (*pb.GetProvinceResponse, error) {
	log.Println("GetProvince method started")

	// Call the API to get the province data
	apiResponse, err := service.CallGetProvinceAPI()
	if err != nil {
		return nil, err
	}

	// Convert the response to the gRPC response format
	provinceList := make([]*pb.Province, len(apiResponse.Rajaongkir.Results))
	for i, result := range apiResponse.Rajaongkir.Results {
		provinceList[i] = &pb.Province{
			ProvinceId:   result.ProvinceID,
			ProvinceName: result.Province,
		}
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

	// Call the API to get the city data
	apiResponse, err := service.CallGetCityAPI(req.ProvinceId)
	if err != nil {
		return nil, err
	}

	// Convert the response to the gRPC response format
	cityList := make([]*pb.City, len(apiResponse.Rajaongkir.Results))
	for i, result := range apiResponse.Rajaongkir.Results {
		cityList[i] = &pb.City{
			CityId:       result.CityID,
			CityName:     result.CityName,
			ProvinceId:   result.ProvinceID,
			ProvinceName: result.Province,
			Type:         result.Type,
			PostalCode:   result.PostalCode,
		}
	}

	// Return the gRPC response
	response := &pb.GetCityResponse{
		Cities: cityList,
	}

	log.Println("GetCity method finished")
	return response, nil
}
