package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitCurrentlyWatchingDependencies() *controllers.CurrentlyWatchingController {
	var dao daoInterfaces.CurrentlyWatchingDaoInterface = daos.NewCurrentlyWatchingDao()
	var service servInterfaces.CurrentlyWatchingServiceInterface = services.NewCurrentlyWatchingService(dao)
	return controllers.NewCurrentlyWatchingController(service)
}
