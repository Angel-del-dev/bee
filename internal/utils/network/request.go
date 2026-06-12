package network

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func ExecuteRequest(payload types.RequestPayload) (types.AgentResponse, error) {
	var agentResponse types.AgentResponse

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return agentResponse, err
	}

	url := "http://127.0.0.1:1234/v1/chat/completions" // TODO Get from configured memory

	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return agentResponse, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return agentResponse, err
	}

	err = json.Unmarshal(body, &agentResponse)
	if err != nil {
		return agentResponse, err
	}

	return agentResponse, nil
}
