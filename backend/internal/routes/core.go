package routes

import (
	"github.com/Roshan-anand/godploy/frontend"
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/Roshan-anand/godploy/internal/middleware"
	authroutes "github.com/Roshan-anand/godploy/internal/routes/auth"
	gitroutes "github.com/Roshan-anand/godploy/internal/routes/git"
	healthroutes "github.com/Roshan-anand/godploy/internal/routes/health"

	projectroutes "github.com/Roshan-anand/godploy/internal/routes/project"
	serviceroutes "github.com/Roshan-anand/godploy/internal/routes/services"
	"github.com/labstack/echo/v5"
)

type Handler struct {
	health  *healthroutes.HealthHandler
	auth    *authroutes.AuthHandler
	service *serviceroutes.ServiceHandler
	git     *gitroutes.GitHandler
	project *projectroutes.ProjectHandler
}

func newHandeler(srv *config.Server) *Handler {
	return &Handler{
		health:  healthroutes.InitHealthHandlers(srv),
		auth:    authroutes.InitAuthHandlers(srv),
		service: serviceroutes.InitServiceHandlers(srv),
		git:     gitroutes.InitGitHandlers(srv),
		project: projectroutes.InitProjectHandlers(srv),
	}

}

// setup all routes
func SetupRoutes(srv *config.Server) (*echo.Echo, error) {
	h := newHandeler(srv)
	m := middleware.NewMiddlewares(srv)
	e := echo.New()

	// initialize static file serving route
	uiFs, err := frontend.GetEmbedFS()
	if err != nil {
		return nil, err
	}
	e.StaticFS("/", uiFs)

	e.Use(m.GlobalMiddlewareCors())

	// health check route
	e.GET("/api/health", func(c *echo.Context) error {
		if srv.DB == nil {
			return c.JSON(500, lib.Res{Message: "database not initialized"})
		}
		return c.JSON(200, lib.Res{Message: "ok"})
	})

	// initialize auth api routes
	authApi := e.Group("/api/auth")
	authApi.GET("/user", h.auth.AuthUser, m.GlobalMiddlewareUser)
	authApi.POST("/register", h.auth.AppRegiter)
	authApi.POST("/login", h.auth.AppLogin)

	// secured routes
	api := e.Group("/api")
	api.Use(m.GlobalMiddlewareUser)

	projectApi := api.Group("/project")
	projectApi.GET("", h.project.GetProjects)
	projectApi.POST("", h.project.CreateProject)
	projectApi.DELETE("", h.project.DeleteProject)

	serviceApi := api.Group("/service")
	serviceApi.POST("/psql", h.service.CreatePsqlService)
	serviceApi.DELETE("/psql", h.service.DeletePsqlService)
	serviceApi.POST("/psql/deploy", h.service.DeployPsqlService)
	serviceApi.POST("/psql/stop", h.service.StopPsqlService)

	githubApi := api.Group("/provider/github")
	githubApi.GET("/app/create", h.git.CreateGithubApp)
	githubApi.GET("/app/callback", h.git.CreateGithubAppCallback)
	githubApi.GET("/app/setup", h.git.SetupGithubApp)
	githubApi.GET("/repo/list", h.git.GetGithubRepoList)

	return e, nil
}
