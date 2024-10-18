package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitSearchDependencies() *controllers.SearchController {
	var dao daoInterfaces.MovieDatabaseDaoInterface = daos.NewTMDBDao()
	var service servInterfaces.SearchServiceInterface = services.NewSearchService(dao)

	return controllers.NewSearchController(service)
}
