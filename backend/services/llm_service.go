package services

import (
	"github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/api"
)

type LLMService struct {
	dao interfaces.LLMDaoInterface
}

func NewLLMService(dao interfaces.LLMDaoInterface) *LLMService {
	return &LLMService{dao}
}

func (service LLMService) AskQuery(messages []api.Message) (string, error) {
	return service.dao.AskQuery(messages)
}
