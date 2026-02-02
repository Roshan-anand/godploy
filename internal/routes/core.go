package routes

import (
	"context"

	ui "github.com/Roshan-anand/godploy/frontend"
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
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
	h := &Handler{Server: srv, Validate: validator.New(), Ctx: context.Background()}
	e := echo.New()

	// initialize static file serving route
	uiFs, err := ui.GetEmbedFS()
	if err != nil {
		return nil, err
	}
	e.StaticFS("/", uiFs)

	// initialize api routes
	api := e.Group("/api")
	api.Use(middleware.CORS("http://localhost:5173"))

	initAuthRoutes(api, h)

	return e, nil
}
