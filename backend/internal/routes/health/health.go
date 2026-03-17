package healthroutes

import (
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/labstack/echo/v5"
)

// to check server health and connectivity with database and other dependencies
//
// route: GET /api/health
func (h *HealthHandler) HealthCheck(c *echo.Context) error {
	if h.Server.DB == nil {
		return c.JSON(500, lib.Res{Message: "database not initialized"})
	}
	return c.JSON(200, lib.Res{Message: "ok"})
}
