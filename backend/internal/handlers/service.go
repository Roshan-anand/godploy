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
// route: GET /api/service/org
func (h *ServiceHandler) GetAllOrganizationServices(c *echo.Context) error {
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)
	q := h.Server.DB.Queries

	orgID, err := q.GetUserCurrentOrg(h.qCtx, u.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "failed to get user's current org"})
	}

	services, err := q.GetAllServicesByOrgId(h.qCtx, orgID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "failed to get services"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"services": services,
	})
}

// get all service deployment jobs
//
// route: GET /api/service/deployment?service_id
func (h *ServiceHandler) GetServiceDeployments(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, lib.Res{Message: "not implemented"})
}

// subscribe to service deployment logs event
//
// route: GET /api/service/deployment/logs?service_id
func (h *ServiceHandler) SubscribeServiceDeploymentLogs(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, lib.Res{Message: "not implemented"})
}
