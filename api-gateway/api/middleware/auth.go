package middleware

import (
	"api-gateway/entity/responses"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
		if authHeader == "" {
			res := &responses.Response{
				Code:    http.StatusUnauthorized,
				Message: "Missing Authorization header",
			}
			return responses.HandleResponse(c, res)
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			res := &responses.Response{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			}
			return responses.HandleResponse(c, res)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		log.Printf("claims: %v", claims)
		if !ok || !token.Valid {
			res := &responses.Response{
				Code:    http.StatusUnauthorized,
				Message: "Invalid Token",
			}
			return responses.HandleResponse(c, res)
		}

		userID, ok := claims["user_id"]
		if !ok {
			res := &responses.Response{
				Code:    http.StatusUnauthorized,
				Message: "Invalid Token",
			}
			return responses.HandleResponse(c, res)
		}
		role, ok := claims["role"]
		if !ok {
			res := &responses.Response{
				Code:    http.StatusUnauthorized,
				Message: "Invalid Token",
			}
			return responses.HandleResponse(c, res)
		}
		c.Set("user_id", userID)
		c.Set("role", role)

		return next(c)
	}
}
