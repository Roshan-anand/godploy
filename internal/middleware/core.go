package middleware

import (
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Middlewares struct {
	Server *config.Server
}

// return new middlewares instance
func NewMiddlewares(s *config.Server) *Middlewares {
	return &Middlewares{Server: s}
}

// global middleware applicable to all routes
func (_ *Middlewares) GlobalMiddleware() echo.MiddlewareFunc {
	return middleware.CORS("http://localhost:5173")
}
