package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitMovieDependencies() *controllers.MovieController {
	var dao daoInterfaces.MovieDatabaseDaoInterface = daos.NewTMDBDao()
	var service servInterfaces.MovieServiceInterface = services.NewMovieService(dao)

	return controllers.NewMovieController(service)
}
