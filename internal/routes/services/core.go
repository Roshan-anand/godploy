package serviceroutes

import (
	"context"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/go-playground/validator/v10"
)

type ServiceHandler struct {
	Server   *config.Server
	Validate *validator.Validate
	Ctx      context.Context
}

func InitServiceHandlers(s *config.Server) *ServiceHandler{
	return &ServiceHandler{
		Server:   s,
		Validate: validator.New(),
		Ctx:      context.Background(),
	}
}