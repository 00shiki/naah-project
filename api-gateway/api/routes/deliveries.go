package routes

import (
	"api-gateway/api/handler/deliveries"
	"api-gateway/api/middleware"
	DELIVERIES_SERVICE "api-gateway/service/deliveries"
	"github.com/labstack/echo/v4"
)

func DeliveryRoute(e *echo.Group, ds DELIVERIES_SERVICE.Service) {
	deliveryHandler := deliveries.NewHandler(ds)
	api := e.Group("/delivery")
	api.Use(middleware.AuthMiddleware)
	api.POST("/cost", deliveryHandler.DeliveryCost)
	api.GET("/provinces", deliveryHandler.ListProvince)
	api.GET("/provinces/:provinceID", deliveryHandler.ListCity)
	api.GET("/couriers", deliveryHandler.ListCourier)
	api.POST("/track", deliveryHandler.InputTrack)
	api.POST("/callback", deliveryHandler.Callback)
}
