package routes

import (
	"context"

	ui "github.com/Roshan-anand/godploy/frontend"
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

const (
	AdminRole  string = "admin"
	MemberRole string = "member"
)

type ErrRes struct {
	Message string `json:"message" validate:"required"`
}

type SuccessRes struct {
	Message string `json:"message" validate:"required"`
}

type Handler struct {
	Server   *config.Server
	Validate *validator.Validate
	Ctx      context.Context
}

// setup all routes
func SetupRoutes(srv *config.Server) (*echo.Echo, error) {
	h := &Handler{Server: srv, Validate: validator.New(), Ctx: context.Background()} // initialize handler
	m := middleware.NewMiddlewares(srv)                                              // initialize middlewares
	e := echo.New()

	// initialize static file serving route
	uiFs, err := ui.GetEmbedFS()
	if err != nil {
		return nil, err
	}
	e.StaticFS("/", uiFs)

	e.Use(m.GlobalMiddlewareCors())

	// initialize auth api routes
	authApi := e.Group("/api/auth")
	authApi.GET("/user", h.authUser, m.GlobalMiddlewareUser)
	authApi.POST("/register", h.appRegiter)
	authApi.POST("/login", h.appLogin)

	// other routes
	// api := e.Group("/api")

	return e, nil
}
