package routes

import (
	"github.com/Roshan-anand/godploy/frontend"
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/handlers"
	"github.com/Roshan-anand/godploy/internal/middleware"
	"github.com/labstack/echo/v5"
)

// setup all routes
func SetupRoutes(srv *config.Server) (*echo.Echo, error) {
	h := handlers.NewHandeler(srv)
	m := middleware.NewMiddlewares(srv)
	e := echo.New()

	// initialize static file serving route
	uiFs, err := frontend.GetEmbedFS()
	if err != nil {
		return nil, err
	}
	e.StaticFS("/", uiFs)

	e.Use(m.GlobalMiddlewareCors())

	api := e.Group("/api")
	public := api.Group("")
	protected := api.Group("")
	protected.Use(m.GlobalMiddlewareUser)

	// public routes
	public.GET("/health", h.Health.HealthCheck)
	public.POST("/url", h.Health.SetUrl)

	// initialize auth api routes
	auth := public.Group("/auth")
	auth.GET("/user", h.Auth.AuthUser, m.GlobalMiddlewareUser)
	auth.POST("/register", h.Auth.AppRegiter)
	auth.POST("/login", h.Auth.AppLogin)

	// initialize org api routes
	org := protected.Group("/org")
	org.GET("", h.Org.GetAllOrgs)
	org.POST("", h.Org.CreateOrg)
	org.DELETE("", h.Org.DeleteOrg)
	org.POST("/switch", h.Org.SwitchOrg)

	// initialize project api routes
	project := protected.Group("/project")
	project.GET("", h.Project.GetProjects)
	project.GET("/all", h.Project.GetProjects)
	project.POST("", h.Project.CreateProject)
	project.DELETE("", h.Project.DeleteProject)

	// initialize service api routes
	service := protected.Group("/service")
	service.POST("/psql", h.Service.CreatePsqlService)
	service.DELETE("/psql", h.Service.DeletePsqlService)
	service.POST("/psql/deploy", h.Service.DeployPsqlService)
	service.POST("/psql/stop", h.Service.StopPsqlService)

	gh := protected.Group("/provider/github")
	ghPublic := public.Group("/provider/github")
	gh.GET("/app/create", h.Git.CreateGithubApp)
	gh.GET("/app", h.Git.GetGithubApp)
	gh.DELETE("/app", h.Git.DeleteGithubApp)
	ghPublic.GET("/app/callback", h.Git.CreateGithubAppCallback)
	ghPublic.GET("/app/setup", h.Git.SetupGithubApp)
	gh.GET("/repo/list", h.Git.GetGithubRepoList)

	return e, nil
}
