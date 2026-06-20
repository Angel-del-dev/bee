package types

type Operation struct {
	ID          string                 `json:"id"`
	Operation   string                 `json:"operation"`
	TargetFile  *string                `json:"target_file"`
	Instruction string                 `json:"instruction"`
	Parameters  map[string]interface{} `json:"parameters"`
}
