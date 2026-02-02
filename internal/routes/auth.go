package routes

import (
	"fmt"
	"net/http"
	"time"

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
	Password string `json:"password" validate:"required,min=8"`
	Org      string `json:"org" validate:"required,min=3,max=50"`
}

// check if user is authenticated
//
// route: GET /api/auth/user
func (h *Handler) authUser(c *echo.Context) error {
	return nil
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

	// check if password length is valid
	switch {
	case len(b.Password) < MIN_PASS_COUNT:
		return c.JSON(http.StatusBadRequest, ErrRes{Message: fmt.Sprintf("Password must be at least %d characters long", MIN_PASS_COUNT)})
	case len(b.Password) > MAX_PASS_COUNT:
		return c.JSON(http.StatusBadRequest, ErrRes{Message: fmt.Sprintf("Password must be at most %d characters long", MAX_PASS_COUNT)})
	}

	// TODO :create seprate table to store creadential like hashed passowrd

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

	sToken, err := lib.GenerateSessionToken()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	// register new admin user
	uId, err := query.AddUser(h.Ctx, db.AddUserParams{
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
	orgId, err := query.InsertOrg(h.Ctx, b.Org)
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

	// store session data
	query.InsertSession(h.Ctx, db.InsertSessionParams{
		UserID:    uId,
		Token:     sToken,
		ExpiresAt: time.Now().Add(lib.SESSION_DATA_EXPIRY_DAY),
	})

	// generate JWT  and setcookie
	jwtStr, err := lib.GenerateJWT(b.Email)
	if err != nil {
		fmt.Println("Generate JWT Error:", err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Internal Server Error"})
	}

	c.SetCookie(&http.Cookie{
		Name:    "godploy_session_data",
		Value:   jwtStr,
		Expires: time.Now().Add(lib.JWT_EXPIRY_HOUR),
	})

	c.SetCookie(&http.Cookie{
		Name:    "godploy_session_token",
		Value:   sToken,
		Expires: time.Now().Add(lib.SESSION_DATA_EXPIRY_DAY),
	})

	return c.JSON(http.StatusOK, SuccessRes{Message: "User Registered Successfully"})
}

// login user
//
// route: POST /api/auth/login
func (h *Handler) appLogin(c *echo.Context) error {
	return nil
}
