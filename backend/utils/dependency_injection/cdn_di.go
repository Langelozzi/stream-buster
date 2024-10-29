package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitCDNDependencies() *controllers.CDNController {
	var dao daoInterfaces.CDNDaoInterface = daos.NewCDNDao()
	var service servInterfaces.CDNServiceInterface = services.NewCDNService(dao)

	return controllers.NewCDNController(service)
}
