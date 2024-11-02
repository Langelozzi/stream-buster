package test

import (
	"testing"

	"github.com/STREAM-BUSTER/stream-buster/daos"
	"github.com/STREAM-BUSTER/stream-buster/models/api"
	"github.com/stretchr/testify/assert"
)

func TestLLMDao(t *testing.T) {

	llmdao := daos.NewLLMDao()

	t.Run("successfully get a response from the llm", func(t *testing.T) {
		messages := []api.Message{
			{Role: "user", Content: "return with the exact following message: hello from ollama"},
		}
		response, err := llmdao.AskQuery(messages)
		assert.Nil(t, err)
		assert.NotNil(t, response)

	})

}
