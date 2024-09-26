package deliveries

import (
	"api-gateway/entity/deliveries"
	pb "api-gateway/proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type DeliveryService struct {
	client pb.DeliveryServiceClient
}

func NewDeliveryService(client pb.DeliveryServiceClient) *DeliveryService {
	return &DeliveryService{
		client: client,
	}
}

func (ds *DeliveryService) DeliveryCost(delivery *deliveries.Delivery) (*pb.DeliveryCostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cartIds := make([]int32, len(delivery.Carts))
	for i, cart := range delivery.Carts {
		cartIds[i] = cart.CartID
	}
	req := &pb.DeliveryCostRequest{
		OriginCityId:      delivery.OriginCityID,
		DestinationCityId: delivery.DestinationCityID,
		CartIds:           cartIds,
		Courier:           delivery.Courier.Name,
	}
	cost, err := ds.client.DeliveryCost(ctx, req)
	if err != nil {
		return nil, err
	}
	return cost, nil
}

func (ds *DeliveryService) GetCouriers() ([]deliveries.Courier, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := ds.client.GetCourier(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	couriers := make([]deliveries.Courier, len(res.Courier))
	for i, c := range res.Courier {
		couriers[i] = deliveries.Courier{
			Name: c,
		}
	}
	return couriers, nil
}

func (ds *DeliveryService) GetProvinces() ([]deliveries.Province, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := ds.client.GetProvince(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	provinces := make([]deliveries.Province, len(res.Provinces))
	for i, p := range res.Provinces {
		provinces[i] = deliveries.Province{
			ID:   p.ProvinceId,
			Name: p.ProvinceName,
		}
	}
	return provinces, nil
}

func (ds *DeliveryService) GetCities(provinceID string) ([]deliveries.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.GetCityRequest{
		ProvinceId: provinceID,
	}
	res, err := ds.client.GetCity(ctx, req)
	if err != nil {
		return nil, err
	}
	cities := make([]deliveries.City, len(res.Cities))
	for i, c := range res.Cities {
		cities[i] = deliveries.City{
			ID:         c.CityId,
			Name:       c.CityName,
			Type:       c.Type,
			PostalCoda: c.PostalCode,
			Province: deliveries.Province{
				ID:   c.ProvinceId,
				Name: c.ProvinceName,
			},
		}
	}
	return cities, nil
}

func (ds *DeliveryService) InputTrackID(orderID int32, trackID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.InputTrackIdRequest{
		OrderId: orderID,
		TrackId: trackID,
	}
	_, err := ds.client.InputTrackId(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DeliveryService) CallbackDelivery(trackID, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.CallbackDeliveryRequest{
		TrackId: trackID,
		Status:  status,
	}
	_, err := ds.client.CallbackDelivery(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
