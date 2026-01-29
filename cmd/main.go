package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/routes"
)

// create and configure the server
func createServer() (*config.Server, error) {
	
	// load server config
	server, err := config.NewServer()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize server: %w", err)
	}

	// setup all routes
	r, err := routes.SetupRoutes(server)
	if err != nil {
		return nil, fmt.Errorf("failed to setup routes: %w", err)
	}

	server.SetupHttp(r) // setup http server with routes

	return server, nil
}

// starts the server
//
// listens for terminate or interrupt signals to shutdown the server gracefully
func runServer(server *config.Server) error {

	// context to listen for terminate or interrupt signals
	notify, cancle := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancle()

	srvErr := make(chan error, 1)
	defer close(srvErr)

	go server.StartServer(srvErr) // start the server

	// graceful shutdown on terminate or interrupt signal
	select {
	case <-notify.Done():
		if err := server.ShutDownServer(); err != nil {
			return err
		}
	case err := <-srvErr:
		return err
	}

	return nil
}

func main() {

	server, err := createServer()
	if err != nil {
		log.Fatal("failed to create server config: ", err)
		return
	}

	if err := runServer(server); err != nil {
		log.Fatal("failed to run server: ", err)
		return
	}

}
