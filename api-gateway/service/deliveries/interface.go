package deliveries

import (
	"api-gateway/entity/deliveries"
	pb "api-gateway/proto"
)

type Service interface {
	DeliveryCost(delivery *deliveries.Delivery) (*pb.DeliveryCostResponse, error)
	GetCouriers() ([]deliveries.Courier, error)
	GetProvinces() ([]deliveries.Province, error)
	GetCities(provinceID string) ([]deliveries.City, error)
	InputTrackID(orderID int32, trackID string) error
	CallbackDelivery(trackID, status string) error
}
