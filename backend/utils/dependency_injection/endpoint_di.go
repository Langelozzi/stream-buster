package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitEndpointDependencies() *controllers.EndpointController {
	var dao daoInterfaces.EndpointDaoInterface = daos.NewEndpointDao()
	var service servInterfaces.EndpointServiceInterface = services.NewEndpointService(dao)

	return controllers.NewEndpointController(service)
}
