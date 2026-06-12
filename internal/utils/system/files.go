package system

import (
	"os"
)

func WorkspacePath() string {
	current_directory, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	return current_directory
}
