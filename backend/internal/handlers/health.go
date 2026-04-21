package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/config"
	deploymentqueue "github.com/Roshan-anand/godploy/internal/jobs/deployment/queue"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type HealthHandler struct {
	Server   *config.Server
	Validate *validator.Validate
	qCtx     context.Context
}

type UrlReq struct {
	Url string `json:"url" validate:"required"`
}

func InitHealthHandlers(s *config.Server) *HealthHandler {
	return &HealthHandler{
		Server:   s,
		Validate: validator.New(),
		qCtx:     context.Background(),
	}
}

// to check server health and connectivity with database and other dependencies
//
// route: GET /api/health
func (h *HealthHandler) HealthCheck(c *echo.Context) error {
	if h.Server.DB == nil {
		return c.JSON(500, lib.Res{Message: "database not initialized"})
	}

	h.Server.DeploymentQ.EnqueuePullJob(&deploymentqueue.PullJobData{})

	return c.JSON(200, lib.Res{Message: "ok"})
}

func (h *HealthHandler) SetUrl(c *echo.Context) error {
	b := new(UrlReq)

	if Res := BindAndValidate(b, c, h.Validate); Res != nil {
		return c.JSON(http.StatusBadRequest, Res)
	}

	fmt.Println("public url :", b.Url)
	h.Server.Config.ServerUrl = b.Url

	return c.JSON(http.StatusOK, nil)
}
