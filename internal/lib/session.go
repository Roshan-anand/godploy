package lib

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

const (
	JWT_EXPIRY_HOUR         = 1 * time.Hour
	JWT_EXPIRY_MIN          = 30 * time.Minute
	SESSION_DATA_EXPIRY_DAY = 7 * 24 * time.Hour
)

// generate JWT token with the given user id
func GenerateJWT(email string) (string, error) {
	secret := []byte("your-256-bit-secret") // TODO : replace with your secret key from env config

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRY_HOUR)),
			Issuer:    "GODPLOY", // TODO : replace with your app name from env config
		},
	})

	jwtStr, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("generate JWT error : %w", err)
	}

	return jwtStr, nil
}

// parse and validate the given JWT token
func VerifyJWT(jwtStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtStr, CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtStr, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse JWT error : %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid JWT token")
	}

	return token.Claims.(*CustomClaims), nil
}

// create a random session token string
func GenerateSessionToken() (string, error) {
	bt := make([]byte, 32)
	if _, err := rand.Read(bt); err != nil {
		return "", fmt.Errorf("generate session token error : %w", err)
	}

	return base64.URLEncoding.EncodeToString(bt), nil
}
