package config

import (
	"database/sql"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/db"
)

type DataBase struct {
	Pool    *sql.DB
	Queries *db.Queries
}

// server holds the global configuration for the application
type Server struct {
	Http *http.Server
	db   *DataBase
	// TODO : add other services like DOCKER client, DB client etc.
}

// creates a new server instance
func NewServer() (*Server, error) {
	// connect DB, Redis, Docker client etc. here and add them to the server struct

	// initialize database connection
	db, err := IntiDb()
	if err != nil {
		return nil, err
	}

	return &Server{
		db: db,
	}, nil
}
