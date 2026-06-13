package types

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Angel-del-dev/bee/internal/utils/system"
)

type Environment struct {
	Host  string `json:"host"`
	Model string `json:"model"`
}

func (e *Environment) Init() {
	agent_path := system.AgentPath()
	configuration_path := filepath.Join(agent_path, "configuration.json")
	if !system.FileExists(configuration_path) {
		return
	}
	// Do not catch errors bellow, if file cannot be read or
	// incorrect parsing happens, both e.Host and e.Model
	// will be empty and will trigger setup dialog

	data, err := os.ReadFile(configuration_path)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, e); err != nil {
		return
	}

}

func (e *Environment) Save() error {
	agentPath := system.AgentPath()
	configurationPath := filepath.Join(agentPath, "configuration.json")

	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(configurationPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
