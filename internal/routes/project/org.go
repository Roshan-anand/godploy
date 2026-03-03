package projectroutes

import (
	"net/http"

	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/labstack/echo/v5"
)

func (h *ProjectHandler) GetOrg(c *echo.Context) error {
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)

	orgs, err := h.Server.DB.Queries.GetAllOrg(h.Ctx, u.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, orgs)
}
