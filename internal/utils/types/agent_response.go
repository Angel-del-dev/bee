package types

type AgentResponseChoiceMessage struct {
	Role             string   `json:"role"`
	Content          string   `json:"content"`
	ReasoningContent string   `json:"reasoning_content"`
	ToolCalls        []string `json:"tool_calls"`
}

type AgentResponseChoice struct {
	Index        int                        `json:"index"`
	Message      AgentResponseChoiceMessage `json:"message"`
	Logprobs     any                        `json:"logprobs"`
	FinishReason string                     `json:"finish_reason"`
}

type AgentResponseUsageDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

type AgentResponseUsage struct {
	PromptTokens            int                       `json:"prompt_tokens"`
	CompletionTokens        int                       `json:"completion_tokens"`
	TotalTokens             int                       `json:"total_tokens"`
	CompletionTokensDetails AgentResponseUsageDetails `json:"completion_tokens_details"`
}

type AgentResponse struct {
	ID                string                `json:"id"`
	Object            string                `json:"object"`
	Created           int                   `json:"created"`
	Model             string                `json:"model"`
	Choices           []AgentResponseChoice `json:"choices"`
	Usage             AgentResponseUsage    `json:"usage"`
	Stats             any                   `json:"stats"`
	SystemFingerprint string                `json:"system_fingerprint"`
}
