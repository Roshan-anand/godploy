package routes

import (
	"fmt"
	"net/http"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/lib"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/moby/moby/api/types/mount"
	"github.com/moby/moby/api/types/swarm"
	"github.com/moby/moby/client"
)

type CreatePsqlServiceReq struct {
	ProjectID   uuid.UUID `json:"project_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	AppName     string    `json:"app_name" validate:"required"`
	Description string    `json:"description"`
	DbName      string    `json:"db_name" validate:"required"`
	DbUser      string    `json:"db_user" validate:"required"`
	DbPassword  string    `json:"db_password" validate:"required"`
	Image       string    `json:"image" validate:"required"`
}

// create a new psql service
//
// route: POST /api/service/psql
func (h *Handler) createPsqlService(c *echo.Context) error {

	b := new(CreatePsqlServiceReq)

	if ErrRes := bindAndValidate(b, c, h.Validate); ErrRes != nil {
		return c.JSON(http.StatusBadRequest, ErrRes)
	}

	// service name to be unique
	b.AppName += lib.GenerateRandomID(6)

	service, err := h.Server.DB.Queries.CreatePsqlService(h.Ctx, db.CreatePsqlServiceParams{
		ID:          lib.NewID(),
		ProjectID:   b.ProjectID,
		ServiceID:   "",
		Name:        b.Name,
		AppName:     b.AppName,
		Description: b.Description,
		DbName:      b.DbName,
		DbUser:      b.DbUser,
		DbPassword:  b.DbPassword, // TODO : make is hased
		Image:       b.Image,
		InternalUrl: "", // TODO : create internal URl
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to create service"})
	}

	return c.JSON(http.StatusOK, service)
}

// deploy the psql service to docker swarm
//
// route: POST /api/service/psql/deploy
func (h *Handler) deployPsqlService(c *echo.Context) error {
	docker := h.Server.Docker
	query := h.Server.DB.Queries

	b := new(ServiceReq)

	if ErrRes := bindAndValidate(b, c, h.Validate); ErrRes != nil {
		return c.JSON(http.StatusBadRequest, ErrRes)
	}

	service, err := query.GetPsqlServiceById(h.Ctx, b.ServiceId)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrRes{Message: "service not found"})
	}

	// create a volume for the service
	vlName := service.AppName + "_pgdata"
	docker.VolumeCreate(h.Ctx, client.VolumeCreateOptions{
		Name:   vlName,
		Driver: "local",
	})

	replicas := uint64(2)

	spec := client.ServiceCreateOptions{
		Spec: swarm.ServiceSpec{

			Annotations: swarm.Annotations{
				Name: service.AppName,
			},

			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image: service.Image,

					Env: []string{
						"POSTGRES_PASSWORD=" + service.DbPassword,
						"POSTGRES_USER=" + service.DbUser,
						"POSTGRES_DB=" + service.DbName,
					},

					Mounts: []mount.Mount{
						{
							Type:   mount.TypeVolume,
							Source: vlName,
							Target: "/var/lib/postgresql/data",
						},
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
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to deploy service"})
	}

	// update the service ID
	if err := query.SetPsqlServiceId(h.Ctx, db.SetPsqlServiceIdParams{
		ServiceID: sRes.ID,
	}); err != nil {
		docker.ServiceRemove(h.Ctx, sRes.ID, client.ServiceRemoveOptions{})
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to update service with service id"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"id": sRes.ID,
	})
}

// stop the psql service
//
// route: POST /api/service/psql/stop
func (h *Handler) stopPsqlService(c *echo.Context) error {
	docker := h.Server.Docker
	query := h.Server.DB.Queries

	b := new(ServiceReq)

	if ErrRes := bindAndValidate(b, c, h.Validate); ErrRes != nil {
		return c.JSON(http.StatusBadRequest, ErrRes)
	}

	service, err := query.GetPsqlServiceById(h.Ctx, b.ServiceId)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrRes{Message: "service not found"})
	}

	fmt.Println("service id :", service.ServiceID)
	if _, err := docker.ServiceRemove(h.Ctx, service.ServiceID, client.ServiceRemoveOptions{}); err != nil {
		fmt.Println("error removing service :", err)
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "error removing service"})
	}

	return c.JSON(http.StatusOK, SuccessRes{Message: "successfully removed the service"})
}

// stops and delete the psql service
//
// route: DELETE /api/service/psql
func (h *Handler) deletePsqlService(c *echo.Context) error {
	docker := h.Server.Docker
	query := h.Server.DB.Queries

	b := new(ServiceReq)

	if ErrRes := bindAndValidate(b, c, h.Validate); ErrRes != nil {
		return c.JSON(http.StatusBadRequest, ErrRes)
	}

	service, err := query.GetPsqlServiceById(h.Ctx, b.ServiceId)
	if err != nil {
		return c.JSON(http.StatusConflict, ErrRes{Message: "Failed to fetch service details"})
	}

	// check and stop the service if it is running
	if s, _ := docker.ServiceInspect(h.Ctx, service.ServiceID, client.ServiceInspectOptions{}); s.Service.ID != "" {
		if _, err := docker.ServiceRemove(h.Ctx, service.ServiceID, client.ServiceRemoveOptions{}); err != nil {
			return c.JSON(http.StatusInternalServerError, ErrRes{Message: fmt.Sprintln("error removing service :", err)})
		}
	}

	if err := query.DeletePsqlService(h.Ctx, b.ServiceId); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRes{Message: "Failed to create service"})
	}

	return c.JSON(http.StatusOK, SuccessRes{
		Message: "Successsfully deleted service",
	})
}
