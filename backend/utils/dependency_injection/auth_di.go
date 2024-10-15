package dependency_injection

import "github.com/DOC-SWAP/Docswap-backend/utils/auth"

func InitAuthDependencies() auth.AuthHandlerInterface {
	var authHandler auth.AuthHandlerInterface = auth.NewAzureAuthHandler()
	return authHandler
}
