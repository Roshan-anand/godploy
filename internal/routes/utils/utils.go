package routeutils

import (
	"fmt"

	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

// binds and validate the given data
func BindAndValidate(b any, c *echo.Context, v *validator.Validate) *lib.Res {

	if err := c.Bind(b); err != nil {
		return &lib.Res{Message: "Invalid Data"}
	}

	if err := v.Struct(b); err != nil {
		return &lib.Res{Message: fmt.Sprintf("validation error : %v", err)}
	}

	return nil
}
