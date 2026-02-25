package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/moby/moby/api/types/swarm"
	"github.com/moby/moby/client"
)

// route: POST /test
func (h *Handler) testRoute(c *echo.Context) error {
	docker := h.Server.Docker
	// query := h.Server.DB.Queries

	replicas := uint64(1)

	spec := client.ServiceCreateOptions{
		Spec: swarm.ServiceSpec{

			Annotations: swarm.Annotations{
				Name: "test-123",
			},

			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image: "postgres:16",

					Env: []string{
						"POSTGRES_PASSWORD=123",
						"POSTGRES_USER=testdb",
						"POSTGRES_DB=testdb",
					},
				},

				RestartPolicy: &swarm.RestartPolicy{
					Condition: swarm.RestartPolicyConditionAny,
				},
			},

			Mode: swarm.ServiceMode{
				Replicated: &swarm.ReplicatedService{
					Replicas: &replicas,
				},
			},
		},
	}

	// depoly the service
	sRes, err := docker.ServiceCreate(h.Ctx, spec)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: fmt.Sprintln("failed to deploy service ", err)})
	}

	return c.JSON(http.StatusOK, sRes)
}
