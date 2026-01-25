package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// server holds the global configuration for the application
type Server struct {
	Http *http.Server
	// TODO : add other services like DOCKER client, DB client etc.
}

func NewAppConfig() *Server {
	return &Server{}
}

// setups new http server with given handler
func (s *Server) SetupHttp(h http.Handler) {
	s.Http = &http.Server{
		Addr:    ":8080", // TODO: change it to env (konf)
		Handler: h,
		// TODO: other configurations
	}
}

// starts the http server
func (s *Server) StartServer(srvErr chan error) {

	if s.Http == nil {
		srvErr <- fmt.Errorf("http server is not initialized")
	}

	// TODO : use logger to log the info

	fmt.Println("starting server on port 8080") // TODO: change it to env (konf)
	if err := s.Http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		srvErr <- err
	}
}

// shuts down the http server gracefully
func (s *Server) ShutDownServer() error {
	ctx, stop := context.WithTimeout(context.Background(), 30*time.Second) // TODO : replace 30 into global variable
	defer stop()

	if err := s.Http.Shutdown(ctx); err != nil {
		// force close the server
		if closeErr := s.Http.Close(); closeErr != nil {
			return errors.Join(err, closeErr)
		}
	}
	return nil
}
