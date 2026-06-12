package services

import (
	"github.com/Roshan-anand/godploy/internal/db"
	deployjob "github.com/Roshan-anand/godploy/internal/jobs/deployment"
	"github.com/Roshan-anand/godploy/internal/jobs/logbroker"
	"github.com/Roshan-anand/godploy/internal/lib/database"
	"github.com/Roshan-anand/godploy/internal/lib/docker"
)

type Services struct {
	Deployment *deployjob.DeploymentService
	LogBroker  *logbroker.LogBrokerService
}

func NewServices(q *db.Queries, docker *docker.DockerClient, badger *database.BadgerDB) *Services {
	logBrokerService := logbroker.NewLogBrokerService(q, badger)
	deploymentService := deployjob.NewDeploymentService(q, docker, logBrokerService)
	return &Services{
		Deployment: deploymentService,
		LogBroker:  logBrokerService,
	}
}
