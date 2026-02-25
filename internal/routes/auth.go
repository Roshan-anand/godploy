package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/Roshan-anand/godploy/internal/types"
	"github.com/labstack/echo/v5"
)

const (
	MAX_PASS_COUNT = 15
	MIN_PASS_COUNT = 8
)

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
	Message string            `json:"message"`
	Name    string            `json:"name"`
	Email   string            `json:"email"`
	Orgs    []db.Organization `json:"orgs"`
}

// check if user is authenticated
//
// route: GET /api/auth/user
func (h *Handler) authUser(c *echo.Context) error {
	u := c.Get(h.Server.Config.EchoCtxUserKey).(lib.AuthUser)

	return c.JSON(http.StatusOK, AuthRes{
		Message: "User Authenticated",
		Name:    u.Name,
		Email:   u.Email,
	})
}

// register a new application
//
// route: POST /api/auth/register
func (h *Handler) appRegiter(c *echo.Context) error {
	b := new(RegisterReq)

	if ErrRes := bindAndValidate(b, c, h.Validate); ErrRes != nil {
		return c.JSON(http.StatusBadRequest, ErrRes)
	}

	query := h.Server.DB.Queries

	// check if admin user already exists
	if exist, err := query.AdminExists(h.Ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	} else if exist {
		return c.JSON(http.StatusBadRequest, ErrRes{Message: "Admin User Already Exists"})
	}

	// hash password
	hPass, err := lib.HashPassword(b.Password)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// register new admin user
	uId, err := query.CreateUser(h.Ctx, db.CreateUserParams{
		ID:       lib.NewID(),
		Name:     b.Name,
		Email:    b.Email,
		HashPass: hPass,
		Role:     types.AdminRole,
	})
	if err != nil {
		fmt.Println("Add User Error:", err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// create organization
	orgId, err := query.CreateOrg(h.Ctx, db.CreateOrgParams{
		ID:   lib.NewID(),
		Name: b.OrgName,
	})
	if err != nil {
		fmt.Println("Insert Org Error:", err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// link user with organization
	if err := query.LinkUserNOrg(h.Ctx, db.LinkUserNOrgParams{
		UserEmail:      b.Email,
		OrganizationID: orgId,
	}); err != nil {
		fmt.Println("Link User N Org Error:", err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// set cookies
	lib.SetSessionCookies(h.Server, c, uId)
	lib.SetJwtCookie(h.Server, c, lib.AuthUser{Email: b.Email, Name: b.Name, Role: types.AdminRole})

	r := AuthRes{
		Message: "Registration Successful",
		Name:    b.Name,
		Email:   b.Email,
		Orgs:    []db.Organization{{ID: orgId, Name: b.OrgName}},
	}
	return c.JSON(http.StatusOK, r)
}

// login user
//
// route: POST /api/auth/login
func (h *Handler) appLogin(c *echo.Context) error {
	b := new(LoginReq)

	if ErrRes := bindAndValidate(b, c, h.Validate); ErrRes != nil {
		return c.JSON(http.StatusBadRequest, ErrRes)
	}

	// get the user
	u, err := h.Server.DB.Queries.GetUserByEmail(h.Ctx, b.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrRes{Message: "user not found"})
	}
	var orgs []db.Organization
	if err := json.Unmarshal([]byte(u.Orgs), &orgs); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// check password
	if !lib.CheckPasswordHash(b.Password, u.HashPass) {
		return c.JSON(http.StatusUnauthorized, ErrRes{Message: "invalid credentials"})
	}

	// set cookies
	lib.SetSessionCookies(h.Server, c, u.ID)
	lib.SetJwtCookie(h.Server, c, lib.AuthUser{Email: u.Email, Name: u.Name, Role: u.Role})

	r := AuthRes{
		Message: "Login Successful",
		Name:    u.Name,
		Email:   u.Email,
		Orgs:    orgs,
	}
	return c.JSON(http.StatusOK, r)
}
