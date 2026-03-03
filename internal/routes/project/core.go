package projectroutes

import (
	"context"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/go-playground/validator/v10"
)

type ProjectHandler struct {
	Server   *config.Server
	Validate *validator.Validate
	Ctx      context.Context
}

func InitProjectHandlers(s *config.Server) *ProjectHandler{
	return &ProjectHandler{
		Server:   s,
		Validate: validator.New(),
		Ctx:      context.Background(),
	}
}