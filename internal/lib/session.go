package lib

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

type AuthUser struct {
	Name  string
	Email string
	Role  types.UserRole
}

type CustomClaims struct {
	jwt.RegisteredClaims
	AuthUser
}

const (
	JWT_EXPIRY_HOUR         = 1 * time.Hour
	JWT_EXPIRY_MIN          = 30 * time.Minute
	SESSION_DATA_EXPIRY_DAY = 7 * 24 * time.Hour
)

// generate JWT token with the given user id
func generateJWT(u AuthUser, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		AuthUser: u,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.Email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRY_HOUR)),
			Issuer:    "GODPLOY", // TODO : replace with your app name from env config
		},
	})

	jwtStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("generate JWT error : %w", err)
	}

	return jwtStr, nil
}

// parse and validate the given JWT token
func VerifyJWT(jwtStr string, secret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
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
func generateSessionToken() (string, error) {
	bt := make([]byte, 32)
	if _, err := rand.Read(bt); err != nil {
		return "", fmt.Errorf("generate session token error : %w", err)
	}

	return base64.URLEncoding.EncodeToString(bt), nil
}

// sets up session token
func SetSessionCookies(s *config.Server, c *echo.Context, uId string) error {
	sToken, err := generateSessionToken()
	if err != nil {
		return fmt.Errorf("generate session token error : %w", err)
	}

	go func() {
		query := s.DB.Queries
		ctx := context.Background()
		// remove old session if any
		if err := query.RemoveSessionByUID(ctx, uId); err != nil {
			fmt.Println("remove old session error :", err)
		}

		// store session data
		if err := query.CreateSession(ctx, db.CreateSessionParams{
			UserID:    uId,
			Token:     sToken,
			ExpiresAt: time.Now().Add(SESSION_DATA_EXPIRY_DAY),
		}); err != nil {
			fmt.Println("create session error :", err)
		}
	}()

	// set session token cookie
	c.SetCookie(&http.Cookie{
		Name:    s.Config.SessionTokenName,
		Value:   sToken,
		Expires: time.Now().Add(SESSION_DATA_EXPIRY_DAY),
		Path:    "/",
	})

	return nil
}

// sets up new JWT cookie
func SetJwtCookie(s *config.Server, c *echo.Context, u AuthUser) error {
	// generate JWT  and setcookie
	jwtStr, err := generateJWT(u, s.Config.JwtSecret)
	if err != nil {
		return fmt.Errorf("generate jwt error : %w", err)
	}

	c.SetCookie(&http.Cookie{
		Name:    s.Config.SessionDataName,
		Value:   jwtStr,
		Expires: time.Now().Add(JWT_EXPIRY_HOUR),
		Path:    "/",
	})

	return nil
}
