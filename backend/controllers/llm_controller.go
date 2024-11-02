package controllers

import (
	"fmt"
	"net/http"

	"github.com/STREAM-BUSTER/stream-buster/models/api"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
)

type LLMController struct {
	service interfaces.LLMServiceInterface
}

func NewLLMController(service interfaces.LLMServiceInterface) *LLMController {
	return &LLMController{
		service: service,
	}
}

func (contr *LLMController) AskQuery(c *gin.Context) {
	var messages []api.Message
	if err := c.BindJSON(&messages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing messages"})
		return
	}

	answer, err := contr.service.AskQuery(messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting response from LLM"})
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, answer)
}
