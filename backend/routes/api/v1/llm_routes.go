package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetLLMRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitLLMDependies()

	group := router.Group("/llm")
	{
		group.POST("/ask-query", controller.AskQuery)
	}
}
