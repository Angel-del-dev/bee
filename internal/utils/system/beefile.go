package system

import (
	"errors"
	"path"
)

func ReadBeeFileAsTask() (string, error) {
	beeFilePath := path.Join(WorkspacePath(), ".bee")
	if !FileExists(beeFilePath) {
		return "", errors.New("Could not find '.bee' file to execute as a taks")
	}

	contents, _ := ReadFileContents(beeFilePath)
	return string(contents), nil
}
