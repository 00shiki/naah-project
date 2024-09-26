package deliveries

import (
	DELIVERIES_PRESENTATION "api-gateway/api/presentation/deliveries"
	"api-gateway/entity/carts"
	"api-gateway/entity/deliveries"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Controller) DeliveryCost(c echo.Context) error {
	payload := new(DELIVERIES_PRESENTATION.DeliveryCostRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	deliveryCarts := make([]carts.CartItem, len(payload.CartIDs))
	for i, id := range payload.CartIDs {
		deliveryCarts[i] = carts.CartItem{
			CartID: id,
		}
	}
	delivery := &deliveries.Delivery{
		OriginCityID:      payload.OriginID,
		DestinationCityID: payload.DestinationID,
		Carts:             deliveryCarts,
		Courier: deliveries.Courier{
			Name: payload.Courier,
		},
	}
	cost, err := handler.ds.DeliveryCost(delivery)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	services := make([]DELIVERIES_PRESENTATION.ServiceItemResponse, len(cost.Service))
	for i, service := range cost.Service {
		services[i] = DELIVERIES_PRESENTATION.ServiceItemResponse{
			ServiceName: service.ServiceName,
			Description: service.Description,
			Cost:        service.Cost,
			Etd:         service.Etd,
		}
	}
	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data: DELIVERIES_PRESENTATION.DeliveryCostResponse{
			Origin: DELIVERIES_PRESENTATION.DeliveryItemResponse{
				CityID:       cost.Origin.CityId,
				CityName:     cost.Origin.CityName,
				ProvinceID:   cost.Origin.ProvinceId,
				ProvinceName: cost.Origin.ProvinceName,
				Type:         cost.Origin.Type,
				PostalCode:   cost.Origin.PostalCode,
			},
			Destination: DELIVERIES_PRESENTATION.DeliveryItemResponse{
				CityID:       cost.Destination.CityId,
				CityName:     cost.Destination.CityName,
				ProvinceID:   cost.Destination.ProvinceId,
				ProvinceName: cost.Destination.ProvinceName,
				Type:         cost.Destination.Type,
				PostalCode:   cost.Destination.PostalCode,
			},
			Services: services,
		},
	}
	return responses.HandleResponse(c, res)
}
