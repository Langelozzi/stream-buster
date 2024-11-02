package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type LLMServiceInterface interface {
	AskQuery(messages []api.Message) (string, error)
}
