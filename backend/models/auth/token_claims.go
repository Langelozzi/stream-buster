package auth

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	Ver      string   `json:"ver"`
	Iss      string   `json:"iss"`
	Sub      string   `json:"sub"`
	Aud      string   `json:"aud"`
	Exp      int64    `json:"exp"`
	Iat      int64    `json:"iat"`
	AuthTime int64    `json:"auth_time"`
	Oid      string   `json:"oid"`
	Emails   []string `json:"emails"`
	NewUser  bool     `json:"newUser"`
	Name     string   `json:"name"`
	Tfp      string   `json:"tfp"`
	Nbf      int64    `json:"nbf"`
	jwt.RegisteredClaims
}
