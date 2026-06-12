package types

type RequestMessagePayload struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestPayload struct {
	Model       string                  `json:"model"`
	Temperature float32                 `json:"temperature"`
	Top_p       float32                 `json:"top_p"`
	Max_tokens  int                     `json:"max_tokens"`
	Stream      bool                    `json:"stream"`
	Messages    []RequestMessagePayload `json:"messages"`
}
