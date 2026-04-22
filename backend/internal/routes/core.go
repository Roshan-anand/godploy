package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Roshan-anand/godploy/frontend"
	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/handlers"
	"github.com/Roshan-anand/godploy/internal/lib/sse"
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

	public.GET("sse", func(c *echo.Context) error {
		log.Printf("SSE client connected, ip: %v", c.RealIP())

		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		count := uint64(0)
		for {
			select {
			case <-c.Request().Context().Done():
				log.Printf("SSE client disconnected, ip: %v", c.RealIP())
				return nil
			case <-ticker.C:
				count++
				event := sse.Event{
					Data: []byte(fmt.Sprintf("count: %d, time: %s\n\n", count, time.Now().Format(time.RFC3339Nano))),
				}
				if err := event.MarshalTo(w); err != nil {
					return err
				}
				if err := http.NewResponseController(w).Flush(); err != nil {
					return err
				}
			}
		}
	})

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
	service.GET("/project", h.Service.GetAllProjectServices)
	service.GET("/org", h.Service.GetAllOrganizationServices)

	psql := service.Group("/psql")
	psql.GET("/:id", h.Service.GetPsqlServiceById)
	psql.POST("", h.Service.CreatePsqlService)
	psql.DELETE("", h.Service.DeletePsqlService)
	psql.POST("/deploy", h.Service.DeployPsqlService)
	psql.POST("/stop", h.Service.StopPsqlService)

	app := service.Group("/app")
	app.GET("/:id", h.Service.GetAppServiceById)
	app.POST("", h.Service.CreateAppService)
	app.DELETE("", h.Service.DeleteAppService)

	gh := protected.Group("/provider/github")
	gh.GET("/app/create", h.Git.CreateGithubApp)
	gh.GET("/app/list", h.Git.GetAllGithubApps)
	gh.DELETE("/app", h.Git.DeleteGithubApp)
	gh.GET("/repo/list", h.Git.GetGithubRepoList)
	ghPublic := public.Group("/provider/github")
	ghPublic.GET("/app/callback", h.Git.CreateGithubAppCallback)
	ghPublic.GET("/app/setup", h.Git.SetupGithubApp)

	return e, nil
}
