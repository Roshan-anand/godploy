package routes

import (
	"net/http"

	"github.com/Roshan-anand/godploy/internal/db"
	// "github.com/Roshan-anand/godploy/internal/lib"
	"github.com/labstack/echo/v5"
)

type CreateProjectReq struct {
	Name           string `json:"name" validate:"required,min=3,max=50"`
	OrganizationID int64  `json:"description" validate:"required"`
}

// create a new project
// 
// route: POST /api/project
func (h *Handler) createProject(c *echo.Context) error {
	// u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)
	b := new(CreateProjectReq)

	errRes := bindAndValidate(b, c, h.Validate)
	if errRes != nil {
		return c.JSON(http.StatusBadRequest, errRes)
	}

	p, err := h.Server.DB.Queries.CreateProject(h.Ctx, db.CreateProjectParams{
		Name:           b.Name,
		OrganizationID: b.OrganizationID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to create project"})
	}

	return c.JSON(http.StatusOK, p)
}
