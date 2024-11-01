package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type LLMDaoInterface interface {
	AskQuery(messages []api.Message) (string, error)
}
