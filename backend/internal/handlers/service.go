package handlers

import (
	"context"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type ServiceHandler struct {
	Server   *config.Server
	Validate *validator.Validate
	qCtx     context.Context
}

func InitServiceHandlers(s *config.Server) *ServiceHandler {
	return &ServiceHandler{
		Server:   s,
		Validate: validator.New(),
		qCtx:     context.Background(),
	}
}

// get all services of a project
//
// route: GET /api/service/project?project_id
func (h *ServiceHandler) GetAllProjectServices(c *echo.Context) error {
	q := h.Server.DB.Queries

	projectId, err := uuid.Parse(c.QueryParam("project_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, lib.Res{Message: "invalid project_id"})
	}

	services, err := q.GetAllServicesByProjectId(h.qCtx, projectId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "failed to get services"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"services": services,
	})
}

// get all services of a organization
//
// route: GET /api/service/org?org_id
func (h *ServiceHandler) GetAllOrganizationServices(c *echo.Context) error {
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)
	q := h.Server.DB.Queries

	orgId, err := uuid.Parse(c.QueryParam("org_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, lib.Res{Message: "invalid org_id"})
	}

	status, Res := CheckUserExistsInOrg(q, u.Email, orgId)
	if Res != nil {
		return c.JSON(status, Res)
	}

	services, err := q.GetAllServicesByOrgId(h.qCtx, orgId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "failed to get services"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"services": services,
	})
}
