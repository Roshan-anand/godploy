package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	SHUTDOWN_TIMEOUT = 30
)

// setups new http server with given handler
//
// @param h : http handler  to set the server with
func (s *Server) SetupHttp(h http.Handler) {
	s.Http = &http.Server{
		Addr:    ":8080", // TODO: change it to env (konf)
		Handler: h,
		// TODO: other configurations
	}
}

// starts the http server
//
// @param srvErr : channel to send server errors
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
	ctx, stop := context.WithTimeout(context.Background(), SHUTDOWN_TIMEOUT*time.Second)
	defer stop()

	if err := s.Http.Shutdown(ctx); err != nil {
		// force close the server
		if closeErr := s.Http.Close(); closeErr != nil {
			return errors.Join(err, closeErr)
		}
	}

	if err := s.CloseDb(); err != nil {
		return err
	}
	return nil
}
