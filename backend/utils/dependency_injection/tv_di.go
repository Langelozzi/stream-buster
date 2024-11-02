package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitTVDependencies() *controllers.TVController {
	var dao daoInterfaces.MovieDatabaseDaoInterface = daos.NewTMDBDao()
	var service servInterfaces.TVServiceInterface = services.NewTVService(dao)

	return controllers.NewTVController(service)
}
