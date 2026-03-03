package routes

import (
	ui "github.com/Roshan-anand/godploy/frontend"
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/Roshan-anand/godploy/internal/middleware"
	authroutes "github.com/Roshan-anand/godploy/internal/routes/auth"
	// gitroutes "github.com/Roshan-anand/godploy/internal/routes/git"
	projectroutes "github.com/Roshan-anand/godploy/internal/routes/project"
	serviceroutes "github.com/Roshan-anand/godploy/internal/routes/services"
	"github.com/labstack/echo/v5"
)

// setup all routes
func SetupRoutes(srv *config.Server) (*echo.Echo, error) {
	authH := authroutes.InitAuthHandlers(srv)
	serviceH := serviceroutes.InitServiceHandlers(srv)
	// gitH := gitroutes.InitGitHandlers(srv)
	projectH := projectroutes.InitProjectHandlers(srv)

	m := middleware.NewMiddlewares(srv) // initialize middlewares
	e := echo.New()

	// initialize static file serving route
	uiFs, err := ui.GetEmbedFS()
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
	authApi.GET("/user", authH.AuthUser, m.GlobalMiddlewareUser)
	authApi.POST("/register", authH.AppRegiter)
	authApi.POST("/login", authH.AppLogin)

	// secured routes
	api := e.Group("/api")
	api.Use(m.GlobalMiddlewareUser)

	projectApi := api.Group("/project")
	projectApi.GET("", projectH.GetProjects)
	projectApi.POST("", projectH.CreateProject)
	projectApi.DELETE("", projectH.DeleteProject)

	serviceApi := api.Group("/service")
	serviceApi.POST("/psql", serviceH.CreatePsqlService)
	serviceApi.DELETE("/psql", serviceH.DeletePsqlService)
	serviceApi.POST("/psql/deploy", serviceH.DeployPsqlService)
	serviceApi.POST("/psql/stop", serviceH.StopPsqlService)

	return e, nil
}
