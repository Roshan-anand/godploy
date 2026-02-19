package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/labstack/echo/v5"
)

type CreateProjectReq struct {
	Name  string `json:"name" validate:"required,min=3,max=50"`
	OrgID int64  `json:"org_id" validate:"required"`
}

type DeleteProjectReq struct {
	Id int64 `json:"name" validate:"required"`
}

// check if user in part of the organization
func CheckUserExistsInOrg(q *db.Queries, email string, orgId int64) (int, *ErrRes) {
	if exists, err := q.CheckUserOrgExists(context.Background(), db.CheckUserOrgExistsParams{
		UserEmail:      email,
		OrganizationID: orgId,
	}); err != nil {
		return http.StatusInternalServerError, &ErrRes{Message: "Failed to create project"}
	} else if !exists {
		return http.StatusForbidden, &ErrRes{Message: "User does not have access to the organization"}
	}

	return http.StatusOK, nil
}

// create a new project
//
// route: POST /api/project
func (h *Handler) createProject(c *echo.Context) error {
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)
	b := new(CreateProjectReq)

	if errRes := bindAndValidate(b, c, h.Validate); errRes != nil {
		return c.JSON(http.StatusBadRequest, errRes)
	}

	q := h.Server.DB.Queries

	status, errRes := CheckUserExistsInOrg(q, u.Email, b.OrgID)
	if errRes != nil {
		return c.JSON(status, errRes)
	}

	// check if project already exists
	if exist, err := q.CheckProjectExist(h.Ctx, db.CheckProjectExistParams{
		OrgID:       b.OrgID,
		ProjectName: b.Name,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "internal server error"})
	} else if exist {
		return c.JSON(http.StatusConflict, ErrRes{Message: fmt.Sprintf("project with name %s already exists ", b.Name)})
	}

	p, err := q.CreateProject(h.Ctx, db.CreateProjectParams{
		Name:           b.Name,
		OrganizationID: b.OrgID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to create project"})
	}

	return c.JSON(http.StatusOK, p)
}

// get all the  projects of the organization
//
// route: GET /api/project?org_id
func (h *Handler) getProjects(c *echo.Context) error {

	// get the value of org_id from query params
	orgIdStr := c.QueryParam("org_id")
	if orgIdStr == "" {
		return c.JSON(http.StatusBadRequest, ErrRes{Message: "org_id query param is required"})
	}

	// convert orgIdStr to int64
	orgId, err := strconv.ParseInt(orgIdStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrRes{Message: "org_id query param must be a valid integer"})
	}

	// TODO : check weather the user exists in the organization or not

	p, err := h.Server.DB.Queries.GetAllProjects(h.Ctx, orgId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to create project"})
	}

	return c.JSON(http.StatusOK, p)
}

// delete a project
//
// route: DELETE /api/project
func (h *Handler) deleteProject(c *echo.Context) error {
	// u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)

	b := new(DeleteProjectReq)

	if ErrRes := bindAndValidate(b, c, h.Validate); ErrRes != nil {
		return c.JSON(http.StatusBadRequest, ErrRes)
	}

	// check if other service exists
	if has, err := h.Server.DB.Queries.CheckProjectHasServices(h.Ctx, b.Id); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to delete project"})
	} else if has {
		return c.JSON(http.StatusConflict, ErrRes{Message: "Project has services associated with it. Please delete the services first."})
	}

	err := h.Server.DB.Queries.DeleteProject(h.Ctx, b.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to create project"})
	}

	return c.JSON(http.StatusOK, SuccessRes{
		Message: "Project deleted successfully",
	})
}
