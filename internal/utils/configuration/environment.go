package configuration

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Angel-del-dev/bee/internal/utils/misc"
	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func LoadEnvironment(scanner *bufio.Scanner) (types.Environment, error) {
	environment := types.Environment{}
	environment.Init()

	if IsEnvironmentValid(environment) {
		return environment, nil
	}

	misc.Speak("Let's configure bee's environment! ")
	misc.Speak("Model host(ip address or domain with port and protocol)")
	fmt.Print("🐝> ")
	if !scanner.Scan() {
		return environment, errors.New("Could not trigger setup, it is required for agent use")
	}

	host := strings.TrimSpace(scanner.Text())

	if host == "" {
		return environment, errors.New("Host cannot be empty, please relaunch the agent")
	}

	misc.Speak("Model's name(Ex: qwen/qwen-4b)")
	fmt.Print("🐝> ")

	if !scanner.Scan() {
		return environment, errors.New("Could not trigger setup, it is required for agent use")
	}

	model := strings.TrimSpace(scanner.Text())

	if model == "" {
		return environment, errors.New("Model cannot be empty, please relaunch the agent")
	}

	err := isAlive(host, model)
	if err != nil {
		return environment, err
	}

	environment.Host = host
	environment.Model = model

	return environment, nil
}

func IsEnvironmentValid(environment types.Environment) bool {
	return strings.TrimSpace(environment.Host) != "" &&
		strings.TrimSpace(environment.Model) != ""
}

func isAlive(host string, model string) error {
	request, err := http.NewRequest("GET", host+"/v1/models", nil)
	if err != nil {
		return errors.New("Invalid host or server not started")
	}
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return errors.New("Server not reachable")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Could not reach server")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.New("failed reading response")
	}

	var models types.ModelsResponse
	err = json.Unmarshal(body, &models)
	if err != nil {
		return errors.New("invalid JSON from server")
	}

	if !modelExists(models, model) {
		return fmt.Errorf("Model %s is not downloaded", model)
	}

	return nil
}

func modelExists(models types.ModelsResponse, target string) bool {
	for _, m := range models.Data {
		if m.ID == target {
			return true
		}
	}
	return false
}
