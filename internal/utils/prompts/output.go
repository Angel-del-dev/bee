package prompts

import (
	"errors"
	"strings"
)

var templateError = errors.New("Could not parse output, try again")

func FormatOutput(output []byte) ([]byte, error) {
	outputString := strings.Trim(string(output), " ")

	firstPosition := strings.IndexByte(outputString, '{')

	if firstPosition < 0 {
		return []byte{}, templateError
	}

	lastPosition := strings.LastIndexByte(outputString, '}')

	if lastPosition < 0 {
		return []byte{}, templateError
	}

	outputLength := (lastPosition - firstPosition) + firstPosition

	return []byte(outputString[firstPosition:(outputLength + 1)]), nil
}
