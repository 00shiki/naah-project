package users

import (
	USERS_PRESENTATION "api-gateway/api/presentation/users"
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (handler *Controller) Detail(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	user, err := handler.us.GetUserDetail(userID)
	if err != nil {
		res := &responses.Response{
			Message: err.Error(),
		}
		switch status.Code(err) {
		case codes.NotFound:
			res.Code = http.StatusNotFound
		case codes.Internal:
			res.Code = http.StatusInternalServerError
		}
		return responses.HandleResponse(c, res)
	}

	res := &responses.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data: USERS_PRESENTATION.DetailResponse{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: user.BirthDate.String(),
			ContactNo: user.ContactNo,
		},
	}
	return responses.HandleResponse(c, res)
}
