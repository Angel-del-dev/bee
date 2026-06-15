package prompts

import (
	"encoding/json"
	"fmt"

	"github.com/Angel-del-dev/bee/internal/utils/system"
)

func Initial(projectTree []system.ProjectNode) string {
	treeJSON, _ := json.MarshalIndent(projectTree, "", " ")
	return fmt.Sprintf(`
You are Bee Planner.

PROJECT TREE:

%s

TASK:

Convert the user request into a JSON array of operations.

CRITICAL RULES:

- Output ONLY JSON.
- The first character of the response must be '['.
- The last character of the response must be ']'.
- No explanations.
- No reasoning.
- No markdown.
- No comments.
- No text before JSON.
- No text after JSON.

Allowed operations:

- read_file
- modify_file
- create_file
- create_documentation
- talk

Operation format:

{
  "id": "op1",
  "operation": "read_file",
  "target_file": "path/to/file",
  "instruction": "task description",
  "parameters": {}
}

Rules:

- Use sequential ids.
- Use project tree paths.
- Prefer read_file before modify_file when file contents are unknown.
- For talk:
  - target_file must be null
  - instruction must be null
  - parameters = {"text":"..."}

VALID:

[
  {
    "id":"op1",
    "operation":"read_file",
    "target_file":"login.py",
    "instruction":"Inspect login implementation and identify logic related to authentication and redirect behaviour.",
    "parameters":{}
  }
]

INVALID:

Here is the plan:

[
  ...
]

INVALID:

I will inspect the login file.

[
  ...
]
  
	`, string(treeJSON))
}
