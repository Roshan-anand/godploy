package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type ProjectHandler struct {
	Server   *config.Server
	Validate *validator.Validate
	qCtx     context.Context
}

type CreateProjectReq struct {
	Name        string    `json:"project_name" validate:"required,min=3"`
	Description string    `json:"description"`
	OrgID       uuid.UUID `json:"org_id" validate:"required"`
}

type DeleteProjectReq struct {
	ID uuid.UUID `json:"id"`
}

func InitProjectHandlers(s *config.Server) *ProjectHandler {
	return &ProjectHandler{
		Server:   s,
		Validate: validator.New(),
		qCtx:     context.Background(),
	}
}

// create a new project
//
// route: POST /api/project
func (h *ProjectHandler) CreateProject(c *echo.Context) error {
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)
	b := new(CreateProjectReq)

	if Res := BindAndValidate(b, c, h.Validate); Res != nil {
		return c.JSON(http.StatusBadRequest, Res)
	}

	q := h.Server.DB.Queries

	status, Res := CheckUserExistsInOrg(q, u.Email, b.OrgID)
	if Res != nil {
		return c.JSON(status, Res)
	}

	// check if project already exists
	if exist, err := q.CheckProjectExist(h.qCtx, db.CheckProjectExistParams{
		OrgID:       b.OrgID,
		ProjectName: b.Name,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, &lib.Res{Message: "internal server error"})
	} else if exist {
		return c.JSON(http.StatusConflict, lib.Res{Message: fmt.Sprintf("project with name %s already exists ", b.Name)})
	}

	p, err := q.CreateProject(h.qCtx, db.CreateProjectParams{
		ID:          lib.NewID(),
		Name:        b.Name,
		Description: b.Description,
		OrgID:       b.OrgID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Failed to create project"})
	}

	return c.JSON(http.StatusOK, p)
}

// get all the  projects of the organization
//
// route: GET /api/project/all?org_id
func (h *ProjectHandler) GetProjects(c *echo.Context) error {
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)

	// get the value of org_id from query params
	orgId, err := uuid.Parse(c.QueryParam("org_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, lib.Res{Message: "invalid organisation id"})
	}
	q := h.Server.DB.Queries

	// TODO : check weather the user exists in the organization or not
	status, Res := CheckUserExistsInOrg(q, u.Email, orgId)
	if Res != nil {
		return c.JSON(status, Res)
	}

	p, err := q.GetAllProjects(h.qCtx, orgId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Failed to get project"})
	}
	return c.JSON(http.StatusOK, p)
}

// delete a project
//
// route: DELETE /api/project
func (h *ProjectHandler) DeleteProject(c *echo.Context) error {
	b := new(DeleteProjectReq)

	if Res := BindAndValidate(b, c, h.Validate); Res != nil {
		return c.JSON(http.StatusBadRequest, Res)
	}

	// TODO : update the query to check other service also when added
	// check if other service exists
	if has, err := h.Server.DB.Queries.CheckProjectHasServices(h.qCtx, b.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Failed to delete project"})
	} else if has {
		return c.JSON(http.StatusConflict, lib.Res{Message: "Project has services associated with it. Please delete the services first."})
	}

	err := h.Server.DB.Queries.DeleteProject(h.qCtx, b.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Failed to delete project"})
	}

	return c.JSON(http.StatusOK, lib.Res{
		Message: "Project deleted successfully",
	})
}
