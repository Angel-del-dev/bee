package agent

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunMainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		fmt.Println("Could not start bee")
		return
	}
	Init()

	for {
		fmt.Print("🐝> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		input = strings.ToLower(input)

		if input == "exit" || input == "quit" {
			fmt.Println("Buzz off 👋")
			break
		}
		if input == "" {
			continue
		}

		fmt.Println("📨 received:", input)
	}
}
