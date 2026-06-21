package operations

import (
	"fmt"
	"regexp"

	"github.com/Angel-del-dev/bee/internal/utils/misc"
	"github.com/Angel-del-dev/bee/internal/utils/prompts"
	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func CreateFile(op types.Operation, memory types.Memory, ctx *types.ProjectContext) error {
	misc.Think(fmt.Sprintf("Preparing '%s'", op.TargetFile))
	prompt := prompts.BuildFilePrompt(op, ctx)
	fmt.Println(prompt)
	response, err := RawCall(prompt, memory)

	if err != nil {
		return err
	}

	re := regexp.MustCompile("(?s)(<!DOCTYPE html>.*)</html>")
	match := re.FindString(response.Choices[0].Message.Content)
	fmt.Println(match)
	ctx.Files[op.TargetFile] = match
	//fmt.Println(content)
	return nil
}
