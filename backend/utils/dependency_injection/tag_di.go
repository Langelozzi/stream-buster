package dependency_injection

import (
	"github.com/DOC-SWAP/Docswap-backend/controllers"
	"github.com/DOC-SWAP/Docswap-backend/daos"
	daoInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/services"
	servInterfaces "github.com/DOC-SWAP/Docswap-backend/services/interfaces"
)

func InitTagDependencies() *controllers.TagController {
	var userTagDao daoInterfaces.UserTagDaoInterface = daos.NewUserTagDao()
	var userTagService servInterfaces.UserTagServiceInterface = services.NewUserTagService(userTagDao)

	var tagDao daoInterfaces.TagDaoInterface = daos.NewTagDao()
	var tagService servInterfaces.TagServiceInterface = services.NewTagService(tagDao, userTagService)

	return controllers.NewTagController(tagService)
}
