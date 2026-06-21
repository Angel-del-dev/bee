package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func RawCall(prompt string, memory types.Memory) (types.AgentResponse, error) {
	var agentResponse types.AgentResponse

	payload := types.RequestPayload{
		Model:       memory.Environment.Model,
		Temperature: 0.2,
		Top_p:       0.8,
		Max_tokens:  1000,
		Messages: []types.RequestMessagePayload{
			{
				Role:    "system",
				Content: prompt,
			},
		},
	}

	url := fmt.Sprintf("%s/v1/chat/completions", memory.Environment.Host)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return agentResponse, err
	}

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

	return agentResponse, err
}
