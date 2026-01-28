package routes

import (
	"context"
	"fmt"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/labstack/echo/v5"
)

// sample route to test database connection
//
// GET /api/test
func (h *Handler) samplerRoute(e *echo.Context) error {
	j := make(map[string]string, 2)

	if err := h.Server.DB.Queries.AddUser(context.Background(), db.AddUserParams{
		Name:     "roshan",
		Email:    "r@gmail.com",
		HashPass: "aasda",
		Role:     "admin",
	}); err != nil {
		fmt.Println("error adding user:", err)
		j["status"] = "failed to add user"
		return e.JSON(500, j)
	}

	j["status"] = "user added successfully"
	return e.JSON(200, j)
}
