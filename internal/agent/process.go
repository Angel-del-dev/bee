package agent

import (
	"fmt"

	"github.com/Angel-del-dev/bee/internal/utils/network"
	"github.com/Angel-del-dev/bee/internal/utils/prompts"
	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func ProcessRequest(request string) {
	payload := types.RequestPayload{
		Model:       "google/gemma-4-e2b", // TODO Get from configured memory
		Temperature: 0.2,
		Top_p:       0.8,
		Max_tokens:  1000,
		Messages: []types.RequestMessagePayload{
			{
				Role:    "system",
				Content: prompts.Initial(),
			},
			{
				Role:    "user",
				Content: request, // TODO add rules based on rules file
			},
		},
	}

	response, err := network.ExecuteRequest(payload)
	if err != nil {
		panic(err) // TODO Proper error handling
	}
	fmt.Print(response)
}
