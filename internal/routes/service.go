package routes

import (
	"github.com/google/uuid"
)

type ServiceReq struct {
	ServiceId uuid.UUID `json:"service_id" validate:"required"`
}
