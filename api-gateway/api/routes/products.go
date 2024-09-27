package routes

import (
	"api-gateway/api/handler/products"
	"api-gateway/api/middleware"
	PRODUCTS_SERVICE "api-gateway/service/products"
	"github.com/labstack/echo/v4"
)

func ProductRoute(e *echo.Group, ps PRODUCTS_SERVICE.Service) {
	productHandler := products.NewHandler(ps)
	api := e.Group("/products")
	api.Use(middleware.AuthMiddleware)
	api.GET("", productHandler.List)
	api.GET("/:productID", productHandler.Detail)
	api.POST("", productHandler.Create)
	api.POST("/:productID", productHandler.CreateDetail)
	api.PUT("/:productID", productHandler.Update)
	api.DELETE("/:productID", productHandler.Delete)
}
