package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/DOC-SWAP/Docswap-backend/models/auth"
	"github.com/DOC-SWAP/Docswap-backend/utils"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"math/big"
	"net/http"
)

type AzureAuthHandler struct{}

func NewAzureAuthHandler() *AzureAuthHandler {
	return &AzureAuthHandler{}
}

// ParseAndValidateToken Azure encrypts the token using RSA, and uses public keys to decrypt the token
// This function fetches the public keys from the JWKS URL and decrypts the token
// We need to fetch the public keys because Azure rotates the keys for security
// The token is then validated for expiration, issuer, audience, etc. based on the standards defined in the OpenID Connect protocol
// https://learn.microsoft.com/en-us/azure/active-directory-b2c/tokens-overview#validation
func (ah *AzureAuthHandler) ParseAndValidateToken(tokenStr string) (*jwt.Token, error) {
	// Fetch JWKs from the URL
	jwksUrl := utils.GetEnvVariable("AZURE_JWKS_URL")
	jwks := fetchJWKS(jwksUrl)

	// Decode the tokenStr from the header
	token, err := jwt.ParseWithClaims(tokenStr, &auth.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid header not found")
		}

		for _, jwk := range jwks.Keys {
			if jwk.Kid == kid {
				return getPublicKeyFromJWK(jwk)
			}
		}

		return nil, fmt.Errorf("public key not found")
	})
	if err != nil {
		return nil, err
	}

	// Get the claims from the token
	claims, ok := token.Claims.(*auth.TokenClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	// Check the iss claim in the token to ensure the token was issued by the correct issuer (Azure AD B2C)
	iss := utils.GetEnvVariable("AZURE_B2C_URL")
	if claims.Iss != iss {
		return nil, fmt.Errorf("invalid issuer")
	}

	// Check the aud claim in the token to ensure the token was issued for the correct application
	aud := utils.GetEnvVariable("AZURE_CLIENT_ID")
	if claims.Aud != aud {
		return nil, fmt.Errorf("invalid audience")
	}

	// Check the iat (issued at) and exp (expiration time) claims in the token to ensure the token is not expired
	currentTime := jwt.TimeFunc().Unix()
	if currentTime < claims.Iat || currentTime > claims.Exp {
		return nil, fmt.Errorf("token is expired")
	}

	return token, nil
}

type JWK struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type JWKs struct {
	Keys []JWK `json:"keys"`
}

// Fetch the JSON Web Key Set (JWKS) from the URL
func fetchJWKS(jwksUrl string) *JWKs {
	// Fetch the JWKS from the URL
	resp, err := http.Get(jwksUrl)
	if err != nil {
		fmt.Println("Error fetching JWKS")
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body")
		}
	}(resp.Body)

	// Decode the JWKS
	var jwks JWKs
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		fmt.Println("Error decoding JWKS")
		return nil
	}

	return &jwks
}

// Retrieve the RSA public key from a JWK
func getPublicKeyFromJWK(jwk JWK) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, err
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, err
	}

	n := new(big.Int).SetBytes(nBytes)

	// Convert eBytes to an integer
	eInt := 0
	for _, b := range eBytes {
		eInt = eInt*256 + int(b)
	}

	pubKey := &rsa.PublicKey{N: n, E: eInt}
	return pubKey, nil
}
