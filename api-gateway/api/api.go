package api

import (
	"api-gateway/api/routes"
	VALIDATOR_PKG "api-gateway/pkg/validator"
	"api-gateway/service/carts"
	"api-gateway/service/orders"
	"api-gateway/service/users"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init(e *echo.Echo, us users.Service, cs carts.Service, os orders.Service) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	e.Use(middleware.Recover())

	e.Validator = &VALIDATOR_PKG.CustomValidator{Validator: validator.New()}

	g := e.Group("/api/v1")
	routes.UserRoute(g, us)
	routes.CartRoute(g, cs)
	routes.OrderRoute(g, os)
}
