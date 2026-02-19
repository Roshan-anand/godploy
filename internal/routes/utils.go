package routes

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

// binds and validate the given data
func bindAndValidate(b any, c *echo.Context, v *validator.Validate) *ErrRes {

	if err := c.Bind(b); err != nil {
		fmt.Println("Bind Error:", err)
		return &ErrRes{Message: "Invalid Data"}
	}

	if err := v.Struct(b); err != nil {
		return &ErrRes{Message: fmt.Sprintf("validation error : %v", err)}
	}

	return nil
}

// 