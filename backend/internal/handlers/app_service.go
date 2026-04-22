package handlers

import (
	"net/http"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/Roshan-anand/godploy/internal/lib/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type CreateAppServiceReq struct {
	ProjectID   uuid.UUID `json:"project_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	AppName     string    `json:"app_name" validate:"required"`
	Description string    `json:"description"`
	GitProvider string    `json:"git_provider" validate:"required"`
	GitRepoID   string    `json:"git_repo_id" validate:"required"`
	GitRepoName string    `json:"git_repo_name" validate:"required"`
	GitBranch   string    `json:"git_branch" validate:"required"`
	BuildPath   string    `json:"build_path" validate:"required"`
}

// create a new app service
//
// route: POST /api/service/app
func (h *ServiceHandler) CreateAppService(c *echo.Context) error {
	b := new(CreateAppServiceReq)

	if Res := BindAndValidate(b, c, h.Validate); Res != nil {
		return c.JSON(http.StatusBadRequest, Res)
	}

	b.AppName += lib.GenerateRandomID(6)

	service, err := h.Server.DB.Queries.CreateAppService(h.qCtx, db.CreateAppServiceParams{
		ID:          lib.NewID(),
		ProjectID:   b.ProjectID,
		Type:        types.AppServiceType,
		ServiceID:   "",
		Name:        b.Name,
		AppName:     b.AppName,
		Description: b.Description,
		GitProvider: b.GitProvider,
		GitRepoID:   b.GitRepoID,
		GitRepoName: b.GitRepoName,
		GitBranch:   b.GitBranch,
		BuildPath:   b.BuildPath,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Failed to create service"})
	}

	return c.JSON(http.StatusOK, service)
}

// get app service details by id
//
// route: GET /api/service/app/:id
func (h *ServiceHandler) GetAppServiceById(c *echo.Context) error {
	serviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, lib.Res{Message: "invalid service id"})
	}

	service, err := h.Server.DB.Queries.GetAppServiceById(h.qCtx, serviceID)
	if err != nil {
		return c.JSON(http.StatusNotFound, lib.Res{Message: "service not found"})
	}

	return c.JSON(http.StatusOK, service)
}

// delete app service
//
// route: DELETE /api/service/app
func (h *ServiceHandler) DeleteAppService(c *echo.Context) error {
	b := new(ServiceReq)

	if Res := BindAndValidate(b, c, h.Validate); Res != nil {
		return c.JSON(http.StatusBadRequest, Res)
	}

	if err := h.Server.DB.Queries.DeleteAppService(h.qCtx, b.ServiceId); err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Failed to delete service"})
	}

	return c.JSON(http.StatusOK, lib.Res{Message: "Successsfully deleted service"})
}
