package routes

import (
	"github.com/labstack/echo/v5"
)

func (h *Handler) samplerRoute(e *echo.Context) error {
	j := map[string]string{"test": "working"}
	return e.JSON(200, j)
}
