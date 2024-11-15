package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models"

type EndpointDaoInterface interface {
	GetAllEndpointsDao() ([]models.Endpoint, error)
}
