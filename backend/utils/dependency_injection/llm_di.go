package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitLLMDependies() *controllers.LLMController {
	var dao daoInterfaces.LLMDaoInterface = daos.NewLLMDao()
	var service servInterfaces.LLMServiceInterface = services.NewLLMService(dao)

	return controllers.NewLLMController(service)
}
