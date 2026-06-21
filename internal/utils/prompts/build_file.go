package prompts

import (
	"fmt"

	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func BuildFilePrompt(op types.Operation, ctx *types.ProjectContext) string {

	context := ""

	for path, content := range ctx.Files {
		context += fmt.Sprintf("\n--- FILE: %s ---\n%s\n", path, content)
	}

	return fmt.Sprintf(`
You are a code generator.

CRITICAL RULES:
- Output ONLY the final file content
- Output EXACTLY ONE version
- Do NOT include reasoning
- Do NOT include multiple attempts
- Do NOT include labels (html, code, etc.)
- Do NOT repeat content
- Do NOT self-correct


PROJECT CONTEXT:
{{%s}}

FILE:
{{%s}}

INSTRUCTION:
{{%s}}
`, context, op.TargetFile, op.Instruction)
}
