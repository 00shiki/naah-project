package orders

import "api-gateway/service/orders"

type Controller struct {
	os orders.Service
}

func NewHandler(os orders.Service) *Controller {
	return &Controller{os}
}
