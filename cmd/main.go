package main

import (
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	e.Static("/", "frontend/dist")
	e.File("/", "frontend/dist/index.html")

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
