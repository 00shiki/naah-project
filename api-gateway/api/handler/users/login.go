package users

import (
	USERS_PRESENTATION "api-gateway/api/presentation/users"
	"api-gateway/entity/responses"
	USERS_ENTITY "api-gateway/entity/users"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (handler *Controller) Login(c echo.Context) error {
	payload := new(USERS_PRESENTATION.LoginRequest)
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

	user := &USERS_ENTITY.User{
		Email:    payload.Email,
		Password: payload.Password,
	}
	token, err := handler.us.LoginUser(user)
	if err != nil {
		res := &responses.Response{
			Message: err.Error(),
		}
		switch status.Code(err) {
		case codes.NotFound:
			res.Code = http.StatusNotFound
		case codes.Unauthenticated:
			res.Code = http.StatusUnauthorized
		case codes.Internal:
			res.Code = http.StatusInternalServerError
		}
		return responses.HandleResponse(c, res)
	}

	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data: USERS_PRESENTATION.LoginResponse{
			Token: *token,
		},
	}
	return responses.HandleResponse(c, res)
}
