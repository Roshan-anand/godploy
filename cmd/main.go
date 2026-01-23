package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/Roshan-anand/godploy/internal/routes"
)

// initialize routes for the server
func initRoutes(srv *config.Server) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hellow server"))
	})

	routes.DockerRoutes(mux, srv.DockerClient)

	srv.HttpSrv.Handler = mux
}

// graceful shutdown
func gracefulShutdown(srv *config.Server) error {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*30)
	defer cancle()

	var gErr error

	// shutdown the server
	if err := srv.HttpSrv.Shutdown(ctx); err != nil {
		fmt.Println("failed to shutdown server:", err)
		fmt.Println("forcefully shutting down the server")
		if forceErr := srv.HttpSrv.Close(); forceErr != nil {
			gErr = errors.Join(err, forceErr)
		}
		gErr = err
	}

	// close the docker client
	if err := srv.DockerClient.Close(); err != nil {
		fmt.Println("failed to close docker client:", err)
	}
	return gErr
}

// runs the server with graceful shutdown
func runServer(srv *config.Server) error {
	notify := make(chan os.Signal, 1)
	srvErr := make(chan error, 1)

	defer func() {
		close(notify)
		close(srvErr)
	}()

	// listen for interrupt  and terminate signals
	signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(notify) // cleanup signal listeners

	// start the server
	go func() {
		fmt.Println("server started on", srv.HttpSrv.Addr)
		if err := srv.HttpSrv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			srvErr <- err
		}
	}()

	select {
	// graceful shutdown
	case <-notify:
		err := gracefulShutdown(srv)
		if err != nil {
			return err
		}

	case err := <-srvErr:
		return err
	}

	return nil
}

func main() {
	srv, err := config.LoadServerConfig()
	if err != nil {
		log.Fatal("failed to load server config:", err)
		return
	}

	initRoutes(srv)
	if err := runServer(srv); err != nil {
		log.Fatal("server error:", err)
	}
}
