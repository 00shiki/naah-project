package products

import "api-gateway/service/products"

type Controller struct {
	ps products.Service
}

func NewHandler(ps products.Service) *Controller {
	return &Controller{ps: ps}
}
