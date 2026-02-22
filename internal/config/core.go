package config

import (
	"database/sql"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/moby/moby/client"
)

type DataBase struct {
	Pool    *sql.DB
	Queries *db.Queries
}

// server holds the global configuration for the application
type Server struct {
	Http   *http.Server
	DB     *DataBase
	Config *Config
	Docker *client.Client
	// TODO : add other services like DOCKER client, DB client etc.
}

// creates a new server instance
func NewServer(cfg *Config) (*Server, error) {
	// connect DB, Redis, Docker client etc. here and add them to the server struct

	// initialize database connection
	db, err := IntiDb(cfg.DbDir)
	if err != nil {
		return nil, err
	}

	//initialize docker client
	docker, err := InitDockerClient()
	if err != nil {
		return nil, err
	}

	return &Server{
		DB:     db,
		Config: cfg,
		Docker: docker,
	}, nil
}
