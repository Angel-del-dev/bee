package system

import (
	"io/fs"
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

func GetProjectTree() []ProjectNode {
	root := WorkspacePath()

	tree := []ProjectNode{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		relPath, _ := filepath.Rel(root, path)
		if relPath == "." {
			return nil
		}

		tree = append(tree, ProjectNode{
			Path:  relPath,
			IsDir: d.IsDir(),
		})
		return nil
	})
	if err != nil {
		panic(err) // TODO Better error handling
	}

	return tree
}

type ProjectNode struct {
	Path  string `json:"path"`
	IsDir bool   `json:"is_dir"`
}
