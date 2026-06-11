package main

import (
	"fmt"
	"os"

	"github.com/Angel-del-dev/bee/internal/agent"
)

func main() {
	current_directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(current_directory)

	agent.Init()
}
