package auth

import "github.com/golang-jwt/jwt/v4"

type AuthHandlerInterface interface {
	ParseAndValidateToken(tokenStr string) (*jwt.Token, error)
}
