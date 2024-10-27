package auth

import "time"

// TokenClaims represents the claims in the JWT token
type TokenClaims struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Issuer    string `json:"iss"`
	ExpiresAt int64  `json:"exp"`
	IssuedAt  int64  `json:"iat"`
}

// NewTokenClaims creates a new TokenClaims with given user data and sets issue/expiration times
func NewTokenClaims(id uint, email, firstName, lastName string) *TokenClaims {
	return &TokenClaims{
		ID:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Issuer:    "auth-service",
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	}
}
