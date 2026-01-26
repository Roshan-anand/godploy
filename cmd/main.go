package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/routes"
)

func main() {
	server := config.NewServer()    // load server config
	r := routes.SetupRoutes(server) // setup all routes
	server.SetupHttp(r)             // setup http server with routes

	// context to listen for terminate or interrupt signals
	sysCtx, cancle := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancle()

	srvErr := make(chan error, 1)
	go server.StartServer(srvErr) // start the server

	select {
	case <-sysCtx.Done():
		if err := server.ShutDownServer(); err != nil {
			log.Fatal("failed to shutdown server: ", err)
		}
	case err := <-srvErr:
		log.Fatal("server error: ", err)
	}

	close(srvErr)
}
