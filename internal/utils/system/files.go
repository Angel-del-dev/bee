package system

import (
	"os"
	"path/filepath"
)

func WorkspacePath() string {
	current_directory, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	return current_directory
}

func AgentPath() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		panic(err)
	}

	return filepath.Dir(exePath)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ReadFileContents(path string) ([]byte, error) {
	contents, err := os.ReadFile(path)
	return contents, err
}
