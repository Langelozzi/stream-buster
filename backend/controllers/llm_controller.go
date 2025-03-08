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

// AskQuery sends a query to the LLM and returns its response
// @Summary Query the LLM
// @Description Sends a list of messages to the LLM and retrieves a response
// @Tags llm
// @Accept  json
// @Produce  json
// @Param messages body []api.Message true "Array of messages for LLM interaction"
// @Success 200 {object} interface{} "Response from the LLM"
// @Failure 400 {object} map[string]interface{} "Error: Invalid request body"
// @Failure 500 {object} map[string]interface{} "Error: Failed to get response from LLM"
// @Router /llm/query [post]
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
