package jwt

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"user_id"`
	Role   int    `json:"role"`
	jwt.RegisteredClaims
}
