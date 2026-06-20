package agent

import (
	"github.com/Angel-del-dev/bee/internal/utils/network"
	"github.com/Angel-del-dev/bee/internal/utils/prompts"
	"github.com/Angel-del-dev/bee/internal/utils/system"
	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func ProcessRequest(request string) (types.AgentResponse, error) {
	projectStructure := system.GetProjectTree()
	payload := types.RequestPayload{
		Model:       Memory.Environment.Model,
		Temperature: 0.2,
		Top_p:       0.8,
		Max_tokens:  1000,
		Messages: []types.RequestMessagePayload{
			{
				Role:    "system",
				Content: prompts.Initial(projectStructure),
			},
			{
				Role:    "user",
				Content: request, // TODO add rules based on rules file
			},
		},
	}
	response, err := network.ExecuteRequest(Memory.Environment.Host, payload)
	if err != nil {
		return types.AgentResponse{}, err
	}
	return response, nil
}
