package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func ExecuteRequest(model_host string, payload types.RequestPayload) (types.AgentResponse, error) {
	var agentResponse types.AgentResponse

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return agentResponse, err
	}

	url := fmt.Sprintf("%s/v1/chat/completions", model_host)

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
