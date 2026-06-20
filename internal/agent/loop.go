package agent

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Angel-del-dev/bee/internal/utils/configuration"
	"github.com/Angel-del-dev/bee/internal/utils/misc"
)

func RunMainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		fmt.Println("Could not start bee")
		return
	}
	Init()
	environment, err := configuration.LoadEnvironment(scanner)
	if err != nil {
		misc.Speak(err.Error())
		return
	}
	err = environment.Save()
	if err != nil {
		misc.Speak(err.Error())
		return
	}

	err = configuration.IsAlive(environment.Host, environment.Model)

	if err != nil {
		misc.Speak(err.Error())
		return
	}

	Memory.Environment = environment

	Init()

	for {
		fmt.Print("🐝> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		input = strings.ToLower(input)

		if input == "exit" {
			misc.Speak("Buzz off 👋")
			break
		}

		workflow, err := ProcessRequest(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		LoopWorkflow(workflow)
	}
}
