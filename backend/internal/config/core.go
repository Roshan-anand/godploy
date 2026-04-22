package config

import (
	"net/http"

	deploymentqueue "github.com/Roshan-anand/godploy/internal/jobs/deployment/queue"
	logbrokerqueue "github.com/Roshan-anand/godploy/internal/jobs/logbroker/queue"
)

// server holds the global configuration for the application
type Server struct {
	Http        *http.Server
	DB          *DataBase
	BadgerDB    *BadgerDB
	Config      *Config
	Docker      *DockerClient
	DeploymentQ *deploymentqueue.JobQueue
	LogBrokerQ  *logbrokerqueue.LogBrokerQueue
}

// creates a new server instance
func NewServer(cfg *Config) (*Server, error) {
	// connect DB, Redis, Docker client etc. here and add them to the server struct

	// initialize database connection
	db, err := InitDb(cfg.DbDir)
	if err != nil {
		return nil, err
	}

	// initialize badgerDB connection
	badger, err := InitBadgerDB(cfg.DbDir)
	if err != nil {
		return nil, err
	}

	//initialize docker client
	docker, err := InitDockerClient()
	if err != nil {
		return nil, err
	}

	// initialize deployment workers queue
	dq := deploymentqueue.InitDeploymentQueue()

	// initialize log broker queue
	lbq := logbrokerqueue.InitLogBrokerQueue()

	return &Server{
		DB:          db,
		BadgerDB:    badger,
		Config:      cfg,
		Docker:      docker,
		DeploymentQ: dq,
		LogBrokerQ:  lbq,
	}, nil
}
