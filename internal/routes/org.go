package routes

import (
	"net/http"

	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/labstack/echo/v5"
)

func (h *Handler) getOrg(c *echo.Context) error {
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)

	orgs, err := h.Server.DB.Queries.GetAllOrg(h.Ctx, u.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{
			Message: "internal server error",
		})
	}
	
	return c.JSON(http.StatusOK, orgs)
}
