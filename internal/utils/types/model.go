package types

type ModelsResponse struct {
	Data []Model `json:"data"`
}

type Model struct {
	ID string `json:"id"`
}
