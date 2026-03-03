package gitroutes

import (
	"context"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/go-playground/validator/v10"
)


type Res struct {
	Message string                 `json:"message" validate:"required"`
	Data    map[string]interface{} `json:"data"`
}

type GitHandler struct {
	Server   *config.Server
	Validate *validator.Validate
	Ctx      context.Context
}

func InitGitHandlers(s *config.Server) *GitHandler {
	return &GitHandler{
		Server:   s,
		Validate: validator.New(),
		Ctx:      context.Background(),
	}
}
