package vouchers

import "api-gateway/service/vouchers"

type Controller struct {
	vs vouchers.Service
}

func NewHandler(vs vouchers.Service) *Controller {
	return &Controller{
		vs: vs,
	}
}
