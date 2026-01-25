package routes

import (
	ui "github.com/Roshan-anand/godploy/frontend"
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Handler struct {
	Server *config.Server
}

// setup all routes
func SetupRoutes(srv *config.Server) *echo.Echo {
	h := &Handler{Server: srv}
	e := echo.New()

	// middlewares
	e.Use(middleware.CORS("http://localhost:5173"))

	// TODO : add some optimizations
	// server static file from /frontend/dist
	e.StaticFS("/", ui.DistDirFS)

	e.GET("/api/test", h.samplerRoute)

	return e
}
