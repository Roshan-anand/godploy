package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/config"
	logbrokerqueue "github.com/Roshan-anand/godploy/internal/jobs/logbroker/queue"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/Roshan-anand/godploy/internal/lib/sse"
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
	q := h.Server.DB.Queries

	serviceID, err := uuid.Parse(c.QueryParam("service_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, lib.Res{Message: "invalid service_id"})
	}

	deployemnts, err := q.GetDeploymentsByServiceID(h.qCtx, serviceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "failed to get deployments"})
	}

	return c.JSON(http.StatusOK, deployemnts)
}

// subscribe to service deployment logs event
//
// route: GET /api/service/deployment/logs?deployment_id
func (h *ServiceHandler) SubscribeServiceDeploymentLogs(c *echo.Context) error {

	deploymentID, err := uuid.Parse(c.QueryParam("deployment_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, lib.Res{Message: "invalid deployment_id"})
	}

	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	sse := sse.NewSSE(w)

	userID := lib.NewID()
	h.Server.LogBrokerQ.SubscribeLogs(userID, &logbrokerqueue.Subscriber{
		SSE:          sse,
		DeploymentID: deploymentID,
	})

	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected, ip: %v", c.RealIP())
			h.Server.LogBrokerQ.UnsubscribeLogs(userID)
			return nil
		}
	}
}
