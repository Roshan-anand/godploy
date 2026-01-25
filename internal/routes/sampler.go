package routes

import "github.com/labstack/echo/v5"

func (h *Handler) samplerRoute(e *echo.Context) error {
	return e.String(200, "")
}
