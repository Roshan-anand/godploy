package routes

import (
	"fmt"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
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
	Org      string `json:"org" validate:"required,min=3,max=50"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}

type AuthRes struct {
	Message string `json:"message"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	// Orgs    []string `json:"orgs,omitempty"`
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

	if err := c.Bind(b); err != nil {
		fmt.Println("Bind Error:", err)
		return c.JSON(http.StatusBadRequest, ErrRes{Message: "Invalid Data"})
	}

	if err := h.Validate.Struct(b); err != nil {
		return c.JSON(http.StatusBadRequest, ErrRes{Message: fmt.Sprintf("validation error : %v", err)})
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
		Name:     b.Name,
		Email:    b.Email,
		HashPass: hPass,
		Role:     AdminRole,
	})
	if err != nil {
		fmt.Println("Add User Error:", err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// create organization
	orgId, err := query.CreateOrg(h.Ctx, b.Org)
	if err != nil {
		fmt.Println("Insert Org Error:", err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// link user with organization
	if err := query.LinkUserNOrg(h.Ctx, db.LinkUserNOrgParams{
		UserID:         uId,
		OrganizationID: orgId,
	}); err != nil {
		fmt.Println("Link User N Org Error:", err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// set cookies
	lib.SetSessionCookies(h.Server, c, uId)
	lib.SetJwtCookie(h.Server, c, lib.AuthUser{Email: b.Email, Name: b.Name})

	r := AuthRes{
		Message: "Registration Successful",
		Name:    b.Name,
		Email:   b.Email,
	}
	return c.JSON(http.StatusOK, r)
}

// login user
//
// route: POST /api/auth/login
func (h *Handler) appLogin(c *echo.Context) error {
	b := new(LoginReq)

	if err := c.Bind(b); err != nil {
		fmt.Println("Bind Error:", err)
		return c.JSON(http.StatusBadRequest, ErrRes{Message: "Invalid Data"})
	}

	if err := h.Validate.Struct(b); err != nil {
		fmt.Println("Validation Error:", err)
		return c.JSON(http.StatusBadRequest, ErrRes{Message: fmt.Sprintf("validation error : %v", err)})
	}

	// get the user
	u, err := h.Server.DB.Queries.GetUserByEmail(h.Ctx, b.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrRes{Message: "user not found"})
	}

	// check password
	if !lib.CheckPasswordHash(b.Password, u.HashPass) {
		return c.JSON(http.StatusUnauthorized, ErrRes{Message: "invalid credentials"})
	}

	// set cookies
	lib.SetSessionCookies(h.Server, c, u.ID)
	lib.SetJwtCookie(h.Server, c, lib.AuthUser{Email: u.Email, Name: u.Name})

	r := AuthRes{
		Message: "Login Successful",
		Name:    u.Name,
		Email:   u.Email,
	}
	return c.JSON(http.StatusOK, r)
}
