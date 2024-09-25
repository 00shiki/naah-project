package routes

import (
	"api-gateway/api/handler/carts"
	"api-gateway/api/middleware"
	CARTS_SERVICE "api-gateway/service/carts"
	"github.com/labstack/echo/v4"
)

func CartRoute(e *echo.Group, cs CARTS_SERVICE.Service) {
	cartHandler := carts.NewHandler(cs)
	api := e.Group("/carts")
	api.Use(middleware.AuthMiddleware)
	api.GET("", cartHandler.ListUserCarts)
	api.POST("", cartHandler.Update)
	api.DELETE("/:cartID", cartHandler.Delete)
}
