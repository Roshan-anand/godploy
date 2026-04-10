package handlers

import "github.com/Roshan-anand/godploy/internal/config"

type Handler struct {
	Health  *HealthHandler
	Auth    *AuthHandler
	Service *ServiceHandler
	Git     *GitHandler
	Project *ProjectHandler
	Org     *OrgHandler
}

func NewHandeler(srv *config.Server) *Handler {
	return &Handler{
		Health:  InitHealthHandlers(srv),
		Auth:    InitAuthHandlers(srv),
		Service: InitServiceHandlers(srv),
		Git:     InitGitHandlers(srv),
		Project: InitProjectHandlers(srv),
		Org:     InitOrgHandlers(srv),
	}
}
