package healthroutes

import (
	"context"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/go-playground/validator/v10"
)

type Res struct {
	Message string                 `json:"message" validate:"required"`
	Data    map[string]interface{} `json:"data"`
}

type HealthHandler struct {
	Server   *config.Server
	Validate *validator.Validate
	qCtx     context.Context
}

func InitHealthHandlers(s *config.Server) *HealthHandler {
	return &HealthHandler{
		Server:   s,
		Validate: validator.New(),
		qCtx:     context.Background(),
	}
}
