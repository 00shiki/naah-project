package users

import (
	USERS_PRESENTATION "api-gateway/api/presentation/users"
	"api-gateway/entity/responses"
	USERS_ENTITY "api-gateway/entity/users"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (handler *Controller) Register(c echo.Context) error {
	payload := new(USERS_PRESENTATION.RegisterRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}
	if err := c.Validate(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	birthDate, err := time.Parse("2006-01-02", payload.BirthDate)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	user := &USERS_ENTITY.User{
		Email:     payload.Email,
		Password:  payload.Password,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		BirthDate: birthDate,
		Address:   payload.Address,
		ContactNo: payload.ContactNo,
		Role:      USERS_ENTITY.Customer,
	}
	if payload.Admin {
		user.Role = USERS_ENTITY.Admin
	}
	err = handler.us.RegisterUser(user)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	res := &responses.Response{
		Code:    http.StatusCreated,
		Message: "Success",
		Data: USERS_PRESENTATION.RegisterResponse{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: user.BirthDate.String(),
			Address:   user.Address,
			ContactNo: user.ContactNo,
		},
	}
	return responses.HandleResponse(c, res)
}
