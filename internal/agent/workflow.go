package agent

import (
	"encoding/json"
	"fmt"

	"github.com/Angel-del-dev/bee/internal/agent/operations"
	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func LoopWorkflow(workflow types.AgentResponse) error {
	queue, err := extractOperations(workflow)
	if err != nil {
		return err
	}

	ctx := &types.ProjectContext{
		Files: make(map[string]string),
	}

	for len(queue) > 0 {
		op := queue[0]
		queue = queue[1:]

		newOps, err := executeOperation(op, ctx)
		if err != nil {
			return err
		}

		queue = append(queue, newOps...)
	}
	fmt.Println(ctx)
	return nil
}

func executeOperation(op types.Operation, ctx *types.ProjectContext) ([]types.Operation, error) {
	switch op.Operation {
	case "create_file":
		return nil, operations.CreateFile(op, Memory, ctx)
	case "talk":
		return nil, operations.Talk(op)
	default:
		return nil, fmt.Errorf("unknown operation: %s", op.Operation)
	}
}

func extractOperations(workflow types.AgentResponse) ([]types.Operation, error) {
	if len(workflow.Choices) == 0 {
		return nil, fmt.Errorf("No choices in response")
	}

	var ops []types.Operation
	err := json.Unmarshal([]byte(workflow.Choices[0].Message.Content), &ops)
	if err != nil {
		return nil, fmt.Errorf("invalid operations JSON: %w", err)
	}

	return ops, nil
}
