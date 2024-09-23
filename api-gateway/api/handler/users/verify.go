package users

import (
	"api-gateway/entity/responses"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"os"
	"strconv"
)

func (handler *Controller) VerifyEmail(c echo.Context) error {
	userIDStr := c.QueryParam("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "user_id must be integer",
		}
		return responses.HandleResponse(c, res)
	}

	err = handler.us.VerifyEmail(int64(userID))
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

	f, err := os.Open("./templates/verified.html")
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}
	body, err := io.ReadAll(f)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}
	return c.HTML(http.StatusOK, string(body))
}
