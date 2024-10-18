package dependency_injection

import "github.com/STREAM-BUSTER/stream-buster/utils/auth"

func InitAuthDependencies() auth.AuthHandlerInterface {
	var authHandler auth.AuthHandlerInterface = auth.NewAzureAuthHandler()
	return authHandler
}
