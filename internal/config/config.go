package config

import (
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	HttpSrv *http.Server
	// TODO : add other configurations like DB, Cache, etc.
}

// loads server configuration from env
func LoadServerConfig() *Server {
	server := &http.Server{
		Addr: os.Getenv("PORT"),
	}

	return &Server{HttpSrv: server}
}
