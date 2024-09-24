package users

import "api-gateway/service/users"

type Controller struct {
	us users.Service
}

func NewHandler(us users.Service) *Controller {
	return &Controller{
		us: us,
	}
}
