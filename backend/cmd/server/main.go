package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/jobs/worker"
	"github.com/Roshan-anand/godploy/internal/routes"
	"github.com/joho/godotenv"
)

// create and configure the server
func createServer() (*config.Server, error) {

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// create server instance
	s, err := config.NewServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize server: %w", err)
	}

	// setup all routes
	r, err := routes.SetupRoutes(s)
	if err != nil {
		return nil, fmt.Errorf("failed to setup routes: %w", err)
	}

	s.SetupHttp(r) // setup http server with routes

	// setup workers
	// TODO : modify the ctx for a gracefull showdown of workers
	w := worker.NewWorker(s)
	go w.PullWorker(context.Background(), s.JobQueue.PullQueue)
	go w.BuildWorker(context.Background(), s.JobQueue.BuildQueue)
	go w.DeployWorker(context.Background(), s.JobQueue.DeployQueue)

	return s, nil
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
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading environment variables from system")
	}

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
