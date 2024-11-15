package interfaces

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
)

type EndpointServiceInterface interface {
	GetAllEndpoints() ([]models.Endpoint, error)
}
