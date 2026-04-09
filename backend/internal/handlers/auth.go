package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/Roshan-anand/godploy/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

const (
	MAX_PASS_COUNT = 15
	MIN_PASS_COUNT = 8
)

type AuthHandler struct {
	Server   *config.Server
	Validate *validator.Validate
	qCtx     context.Context
}

type RegisterReq struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15"`
	OrgName  string `json:"org_name" validate:"required,min=3,max=50"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}

type AuthRes struct {
	Message string    `json:"message"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	OrgId   uuid.UUID `json:"org_id"`
	OrgName string    `json:"org_name"`
}

func InitAuthHandlers(s *config.Server) *AuthHandler {
	return &AuthHandler{
		Server:   s,
		Validate: validator.New(),
		qCtx:     context.Background(),
	}
}

// check if user is authenticated
//
// route: GET /api/auth/user
func (h *AuthHandler) AuthUser(c *echo.Context) error {
	u, ok := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)

	if !ok {
		exists, err := h.Server.DB.Queries.AdminExists(h.qCtx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Internal Sever Error"})
		}

		if !exists {
			return c.JSON(http.StatusForbidden, lib.Res{Message: "No admin registered"})
		}
		return c.JSON(http.StatusUnauthorized, lib.Res{Message: "Unauthorized"})
	}

	return c.JSON(http.StatusOK, AuthRes{
		Message: "User Authenticated",
		Name:    u.Name,
		Email:   u.Email,
		OrgId:   u.OrgId,
		OrgName: u.OrgName,
	})
}

// register a new application
//
// route: POST /api/auth/register
func (h *AuthHandler) AppRegiter(c *echo.Context) error {
	b := new(RegisterReq)

	if Res := BindAndValidate(b, c, h.Validate); Res != nil {
		return c.JSON(http.StatusBadRequest, Res)
	}

	query := h.Server.DB.Queries

	// check if admin user already exists
	if exist, err := query.AdminExists(h.qCtx); err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Internal Server Error"})
	} else if exist {
		return c.JSON(http.StatusBadRequest, lib.Res{Message: "Admin User Already Exists"})
	}

	// hash password
	hPass, err := lib.HashPassword(b.Password)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Internal Server Error"})
	}

	// create organization first (user needs orgId at insert time)
	orgId, err := query.CreateOrg(h.qCtx, db.CreateOrgParams{
		ID:   lib.NewID(),
		Name: b.OrgName,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Internal Server Error"})
	}

	// register new admin user
	uId, err := query.CreateUser(h.qCtx, db.CreateUserParams{
		ID:           lib.NewID(),
		Name:         b.Name,
		Email:        b.Email,
		HashPass:     hPass,
		Role:         types.AdminRole,
		CurrentOrgID: orgId,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Internal Server Error"})
	}

	// link user with organization
	if err := query.LinkUserNOrg(h.qCtx, db.LinkUserNOrgParams{
		UserEmail:      b.Email,
		OrganizationID: orgId,
	}); err != nil {
		fmt.Println("Link User N Org Error:", err)
		return c.JSON(http.StatusInternalServerError, lib.Res{Message: "Internal Server Error"})
	}

	// set cookies
	lib.SetSessionCookies(h.Server, c, uId)
	lib.SetJwtCookie(h.Server, c, lib.AuthUser{Email: b.Email, Name: b.Name, Role: types.AdminRole})

	r := AuthRes{
		Message: "Registration Successful",
		Name:    b.Name,
		Email:   b.Email,
		OrgId:   orgId,
		OrgName: b.OrgName,
	}
	return c.JSON(http.StatusOK, r)
}

// login user
//
// route: POST /api/auth/login
func (h *AuthHandler) AppLogin(c *echo.Context) error {
	b := new(LoginReq)

	if Res := BindAndValidate(b, c, h.Validate); Res != nil {
		return c.JSON(http.StatusBadRequest, Res)
	}

	// get the user
	u, err := h.Server.DB.Queries.GetUserByEmail(h.qCtx, b.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, lib.Res{Message: "user not found"})
	}

	// check password
	if !lib.CheckPasswordHash(b.Password, u.HashPass) {
		return c.JSON(http.StatusUnauthorized, lib.Res{Message: "invalid credentials"})
	}

	// set cookies
	lib.SetSessionCookies(h.Server, c, u.ID)
	lib.SetJwtCookie(h.Server, c, lib.AuthUser{Email: u.Email, Name: u.Name, Role: u.Role})

	r := AuthRes{
		Message: "Login Successful",
		Name:    u.Name,
		Email:   u.Email,
		OrgId:   u.OrgID,
		OrgName: u.OrgName,
	}
	return c.JSON(http.StatusOK, r)
}
