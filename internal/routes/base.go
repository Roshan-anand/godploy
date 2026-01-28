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
func SetupRoutes(srv *config.Server) (*echo.Echo, error) {
	h := &Handler{Server: srv}
	e := echo.New()

	uiFs, err := ui.GetEmbedFS() // get embedded frontend filesystem
	if err != nil {
		return nil, err
	}
	e.StaticFS("/", uiFs)

	api := e.Group("/api")
	api.Use(middleware.CORS("http://localhost:5173"))

	api.GET("/test", h.samplerRoute)

	return e, nil
}
