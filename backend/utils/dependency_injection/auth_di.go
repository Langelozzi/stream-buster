package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/daos"
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	iServices "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitAuthDependencies() iServices.AuthServiceInterface {
	var authDao iDao.AuthDaoInterface = daos.NewAuthDao()
	var authService iServices.AuthServiceInterface = services.NewAuthService(authDao)
	return authService
}
