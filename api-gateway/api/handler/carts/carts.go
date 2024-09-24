package carts

import "api-gateway/service/carts"

type Controller struct {
	cs carts.Service
}

func NewHandler(cs carts.Service) *Controller {
	return &Controller{
		cs: cs,
	}
}
