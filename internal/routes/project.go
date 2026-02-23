package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/labstack/echo/v5"
)

type DeleteProjectReq struct {
	ID string `json:"id"`
}

// check if user in part of the organization
func CheckUserExistsInOrg(q *db.Queries, email string, orgId string) (int, *ErrRes) {
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
	b := new(db.CreateProjectParams)

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
		Name:  b.Name,
		OrgID: b.OrgID,
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
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)

	// get the value of org_id from query params
	orgId := c.QueryParam("org_id")
	if orgId == "" {
		return c.JSON(http.StatusBadRequest, ErrRes{Message: "invalid organisation id"})
	}
	q := h.Server.DB.Queries

	// TODO : check weather the user exists in the organization or not
	status, errRes := CheckUserExistsInOrg(q, u.Email, orgId)
	if errRes != nil {
		return c.JSON(status, errRes)
	}

	p, err := q.GetAllProjects(h.Ctx, orgId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to get project"})
	}

	return c.JSON(http.StatusOK, p)
}

// delete a project
//
// route: DELETE /api/project
func (h *Handler) deleteProject(c *echo.Context) error {
	b := new(DeleteProjectReq)

	if ErrRes := bindAndValidate(b, c, h.Validate); ErrRes != nil {
		return c.JSON(http.StatusBadRequest, ErrRes)
	}

	// check if other service exists
	if has, err := h.Server.DB.Queries.CheckProjectHasServices(h.Ctx, b.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to delete project"})
	} else if has {
		return c.JSON(http.StatusConflict, ErrRes{Message: "Project has services associated with it. Please delete the services first."})
	}

	err := h.Server.DB.Queries.DeleteProject(h.Ctx, b.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to delete project"})
	}

	return c.JSON(http.StatusOK, SuccessRes{
		Message: "Project deleted successfully",
	})
}
