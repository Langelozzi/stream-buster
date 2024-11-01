package daos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STREAM-BUSTER/stream-buster/models/api"
	"github.com/STREAM-BUSTER/stream-buster/utils"
)

type LLMDao struct{}

func NewLLMDao() *LLMDao {
	return &LLMDao{}
}

// AskQuery sends a request to the /ask-query endpoint with the given messages
func (LLMDao) AskQuery(messages []api.Message) (string, error) {
	url := fmt.Sprintf("%s/ask-query", utils.GetEnvVariable("LLM_BACKEND_URL"))
	requestPayload := map[string]interface{}{
		"messages": messages,
	}

	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		return "", fmt.Errorf("error marshalling request payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating new request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", utils.GetEnvVariable("LLM_BACKEND_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Reply string `json:"reply"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	return response.Reply, nil
}
