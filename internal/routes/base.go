package routes

import (
	ui "github.com/Roshan-anand/godploy/frontend"
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/labstack/echo/v5"
)

type Handler struct {
	Server *config.Server
}

// setup all routes
func SetupRoutes(srv *config.Server) *echo.Echo {
	// h := &Handler{Server: srv}
	e := echo.New()

	// TODO : add some optimizations
	// server static file from /frontend/dist
	e.StaticFS("/", ui.DistDirFS)

	return e
}
