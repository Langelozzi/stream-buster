package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models"
)

type EndpointService struct {
	dao iDao.EndpointDaoInterface
}

func NewEndpointService(dao iDao.EndpointDaoInterface) *EndpointService {
	return &EndpointService{dao: dao}
}

func (service *EndpointService) GetAllEndpoints() ([]models.Endpoint, error) {
	return service.dao.GetAllEndpointsDao()
}
