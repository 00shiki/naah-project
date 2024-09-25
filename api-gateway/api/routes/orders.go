package routes

import (
	"api-gateway/api/handler/orders"
	ORDERS_SERVICE "api-gateway/service/orders"
	"github.com/labstack/echo/v4"
)

func OrderRoute(e *echo.Group, os ORDERS_SERVICE.Service) {
	orderHandler := orders.NewHandler(os)
	api := e.Group("/orders")
	api.POST("/callback", orderHandler.Callback)
	protected := api.Group("/protected")
	protected.POST("", orderHandler.Create)
	protected.GET("", orderHandler.List)
}
