package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitMediaDependencies() *controllers.MediaController {
	var cdnDao daoInterfaces.CDNDaoInterface = daos.NewCDNDao()
	var cdnService servInterfaces.CDNServiceInterface = services.NewCDNService(cdnDao)

	var dao daoInterfaces.MediaDaoInterface = daos.NewMediaDao()
	var service servInterfaces.MediaServiceInterface = services.NewMediaService(dao, cdnService)

	return controllers.NewMediaController(service)
}
