package agent

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Angel-del-dev/bee/internal/utils/configuration"
)

func RunMainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		fmt.Println("Could not start bee")
		return
	}
	Init()
	configuration, err := configuration.LoadEnvironment(scanner)
	if err != nil {
		fmt.Println("🐝 ", err.Error())
		return
	}
	err = configuration.Save()
	if err != nil {
		fmt.Println("🐝 ", err.Error())
		return
	}

	Memory.Environment = configuration

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
			fmt.Println("Buzz off 👋")
			break
		}

		ProcessRequest(input)
	}
}
