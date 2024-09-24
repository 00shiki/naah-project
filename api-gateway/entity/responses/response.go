package responses

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Response struct {
	Code    int         `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func HandleResponse(c echo.Context, res *Response) error {
	if res.Code == http.StatusInternalServerError {
		log.Printf("%s: %v", c.Request().URL, res.Message)
		res.Message = "Internal Server Error"
		return c.JSON(res.Code, res)
	}
	return c.JSON(res.Code, res)
}
