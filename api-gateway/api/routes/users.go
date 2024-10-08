package routes

import (
	"api-gateway/api/handler/users"
	"api-gateway/api/middleware"
	USERS_SERVICE "api-gateway/service/users"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Group, us USERS_SERVICE.Service) {
	userHandler := users.NewHandler(us)
	api := e.Group("/users")
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/verify", userHandler.VerifyEmail)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware)
	protected.GET("", userHandler.Detail)
}
