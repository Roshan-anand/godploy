package config

import (
	"fmt"
	"net/http"
	"os"

	"github.com/docker/docker/client"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	HttpSrv      *http.Server
	DockerClient *client.Client
	// TODO : add other configurations like DB, Cache, etc.
}

// loads server configuration from env
func LoadServerConfig() (*Server, error) {
	server := &http.Server{
		Addr: os.Getenv("PORT"),
	}

	client, err := getDockerClient()
	if err != nil {
		return nil, fmt.Errorf("Docker unable to connect : %w", err)
	}

	return &Server{HttpSrv: server, DockerClient: client}, nil
}
