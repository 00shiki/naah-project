package handler

import (
	"context"
	"database/sql"
	"log"
	"order-service/pb"
	"order-service/service"
	"order-service/utils"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DeliveryHandler struct {
	db  *sql.DB
	rmq *utils.RabbitMQClient
}

func NewDeliveryHandler(db *sql.DB, rmq *utils.RabbitMQClient) *DeliveryHandler {
	return &DeliveryHandler{db, rmq}
}

// DeliveryCost method
func (h *DeliveryHandler) DeliveryCost(ctx context.Context, req *pb.DeliveryCostRequest) (*pb.DeliveryCostResponse, error) {
	log.Println("DeliveryCost method started")

	cartIds := utils.RemoveDuplicates(req.CartIds)

	// Calculate total weight
	var weight int32
	for _, cartID := range cartIds {
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
	services := make([]*pb.ServiceItem, len(apiResponse.Rajaongkir.Results[0].Costs))

	for i, cost := range apiResponse.Rajaongkir.Results[0].Costs {
		services[i] = &pb.ServiceItem{
			ServiceName: cost.Service,
			Description: cost.Description,
			Cost:        int32(cost.Cost[0].Value),
			Etd:         cost.Cost[0].Etd,
		}
	}

	response := &pb.DeliveryCostResponse{
		Origin:      &origin,
		Destination: &destination,
		Service:     services,
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

func (h *DeliveryHandler) GetCourier(ctx context.Context, req *empty.Empty) (*pb.GetCourierResponse, error) {

	response := &pb.GetCourierResponse{
		Courier: []string{"jne", "pos", "tiki"},
	}
	return response, nil
}

// CallbackDelivery handles updating delivery based on callback information
func (h *DeliveryHandler) CallbackDelivery(ctx context.Context, req *pb.CallbackDeliveryRequest) (*emptypb.Empty, error) {
	log.Printf("CallbackDelivery initiated with Track ID: %s, Status: %s\n", req.TrackId, req.Status)

	// Start a transaction
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error starting transaction: %v\n", err)
		return nil, status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	}

	// If status is 'delivered', update both order and delivery statuses
	if strings.EqualFold(req.Status, "delivered") {
		// Update order status to 'delivered'
		updateOrderQuery := `UPDATE orders SET status = 'delivered' 
							 WHERE order_id = (SELECT order_id FROM deliveries WHERE track_id = ?)`
		log.Printf("Executing query to update order status: %s\n", updateOrderQuery)

		orderResult, err := tx.ExecContext(ctx, updateOrderQuery, req.TrackId)
		if err != nil {
			tx.Rollback()
			log.Printf("Error updating order status for Track ID %s: %v\n", req.TrackId, err)
			return nil, status.Errorf(codes.Internal, "failed to update order status: %v", err)
		}

		rowsAffected, err := orderResult.RowsAffected()
		if err != nil || rowsAffected == 0 {
			tx.Rollback()
			log.Printf("No order found for Track ID: %s\n", req.TrackId)
			return nil, status.Errorf(codes.NotFound, "no order found for the provided track ID: %s", req.TrackId)
		}

		// Update delivery status and set the arrival date
		updateDeliveryQuery := `UPDATE deliveries 
								SET status = ?, arrival_date = CURRENT_TIMESTAMP 
								WHERE track_id = ?`
		log.Printf("Executing query to update delivery status: %s\n", updateDeliveryQuery)

		deliveryResult, err := tx.ExecContext(ctx, updateDeliveryQuery, req.Status, req.TrackId)
		if err != nil {
			tx.Rollback()
			log.Printf("Error updating delivery status for Track ID %s: %v\n", req.TrackId, err)
			return nil, status.Errorf(codes.Internal, "failed to update delivery status: %v", err)
		}

		rowsAffected, err = deliveryResult.RowsAffected()
		if err != nil || rowsAffected == 0 {
			tx.Rollback()
			log.Printf("No delivery found for Track ID: %s\n", req.TrackId)
			return nil, status.Errorf(codes.NotFound, "no delivery found for the provided track ID: %s", req.TrackId)
		}

		// Retrieve the user email and order ID to send notification
		var email string
		var orderId int
		getEmailQuery := `SELECT users.email, orders.order_id 
						  FROM deliveries
						  JOIN orders ON deliveries.order_id = orders.order_id
						  JOIN users ON orders.user_id = users.user_id
						  WHERE deliveries.track_id = ?`
		log.Printf("Executing query to fetch email and order ID: %s\n", getEmailQuery)

		err = tx.QueryRowContext(ctx, getEmailQuery, req.TrackId).Scan(&email, &orderId)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				log.Printf("No email or order found for Track ID: %s\n", req.TrackId)
				return nil, status.Errorf(codes.NotFound, "no associated user found for the provided track ID: %v", err)
			}
			log.Printf("Error fetching email and order ID for Track ID %s: %v\n", req.TrackId, err)
			return nil, status.Errorf(codes.Internal, "error retrieving associated email or order: %v", err)
		}

		// Send email notification for successful delivery
		log.Printf("Sending delivered notification email to: %s for Order ID: %d\n", email, orderId)
		utils.SendDeliveredEmail(email, orderId, h.rmq)

	} else {
		// If status is not 'delivered', just update the delivery status
		updateDeliveryQuery := `UPDATE deliveries SET status = ? WHERE track_id = ?`
		log.Printf("Executing query to update delivery status for non-delivered status: %s\n", updateDeliveryQuery)

		deliveryResult, err := tx.ExecContext(ctx, updateDeliveryQuery, req.Status, req.TrackId)
		if err != nil {
			tx.Rollback()
			log.Printf("Error updating delivery status for Track ID %s: %v\n", req.TrackId, err)
			return nil, status.Errorf(codes.Internal, "failed to update delivery status: %v", err)
		}

		rowsAffected, err := deliveryResult.RowsAffected()
		if err != nil || rowsAffected == 0 {
			tx.Rollback()
			log.Printf("No delivery found for Track ID: %s\n", req.TrackId)
			return nil, status.Errorf(codes.NotFound, "no delivery found for the provided track ID: %s", req.TrackId)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction for Track ID %s: %v\n", req.TrackId, err)
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}

	log.Printf("Successfully updated delivery and order status for Track ID: %s, Status: %s\n", req.TrackId, req.Status)
	return &emptypb.Empty{}, nil
}

func (h *DeliveryHandler) InputTrackId(ctx context.Context, req *pb.InputTrackIdRequest) (*emptypb.Empty, error) {
	// Begin transaction
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error starting transaction: %v\n", err)
		return nil, status.Errorf(codes.Internal, "error starting transaction")
	}

	// Check the current status in the orders table
	var currentStatus string
	query := `SELECT status FROM orders WHERE order_id = ?`
	err = tx.QueryRowContext(ctx, query, req.OrderId).Scan(&currentStatus)
	if err != nil {
		tx.Rollback()
		log.Printf("Error retrieving order status for Order ID %d: %v\n", req.OrderId, err)
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "order not found for ID: %d", req.OrderId)
		}
		return nil, status.Errorf(codes.Internal, "error retrieving order status: %v", err)
	}

	// Update the order status only if the current status is "PAID" (case-insensitive)
	if strings.EqualFold(currentStatus, "PAID") {
		// Update the delivery with the given order_id and set the track_id and status
		query = `UPDATE deliveries SET track_id = ?, status = ?, delivery_date = CURRENT_TIMESTAMP WHERE order_id = ?`
		result, err := tx.ExecContext(ctx, query, req.TrackId, "manifested", req.OrderId)
		if err != nil {
			tx.Rollback()
			log.Printf("Error updating delivery for Order ID %d: %v\n", req.OrderId, err)
			return nil, status.Errorf(codes.Internal, "error updating delivery: %v", err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil || rowsAffected == 0 {
			tx.Rollback()
			log.Printf("No delivery found for order_id: %d\n", req.OrderId)
			return nil, status.Errorf(codes.NotFound, "no delivery found for the given order_id: %d", req.OrderId)
		}

		// Update the delivery with the given order_id and set the track_id and status
		query = `UPDATE orders SET status = ? WHERE order_id = ?`
		result, err = tx.ExecContext(ctx, query, "manifested", req.OrderId)
		if err != nil {
			tx.Rollback()
			log.Printf("Error updating delivery for Order ID %d: %v\n", req.OrderId, err)
			return nil, status.Errorf(codes.Internal, "error updating delivery: %v", err)
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil || rowsAffected == 0 {
			tx.Rollback()
			log.Printf("No delivery found for order_id: %d\n", req.OrderId)
			return nil, status.Errorf(codes.NotFound, "no order found for the given order_id: %d", req.OrderId)
		}

	} else {
		tx.Rollback()
		log.Printf("Order ID %d has not yet been paid\n", req.OrderId)
		return nil, status.Errorf(codes.InvalidArgument, "Order ID %d not yet paid", req.OrderId)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction for Order ID %d: %v\n", req.OrderId, err)
		return nil, status.Errorf(codes.Internal, "error committing transaction: %v", err)
	}

	log.Printf("Successfully updated delivery for Order ID: %d with Track ID: %s\n", req.OrderId, req.TrackId)
	return &emptypb.Empty{}, nil
}
