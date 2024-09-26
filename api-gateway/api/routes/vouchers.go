package routes

import (
	"api-gateway/api/handler/vouchers"
	"api-gateway/api/middleware"
	VOUCHERS_SERVICE "api-gateway/service/vouchers"
	"github.com/labstack/echo/v4"
)

func VoucherRoute(e *echo.Group, vs VOUCHERS_SERVICE.Service) {
	voucherHandler := vouchers.NewHandler(vs)
	api := e.Group("/vouchers")
	api.Use(middleware.AuthMiddleware)
	api.GET("", voucherHandler.List)
	api.GET("/:voucherID", voucherHandler.Detail)
	api.POST("", voucherHandler.Create)
}
