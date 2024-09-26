package deliveries

import "api-gateway/service/deliveries"

type Controller struct {
	ds deliveries.Service
}

func NewHandler(ds deliveries.Service) *Controller {
	return &Controller{ds: ds}
}
