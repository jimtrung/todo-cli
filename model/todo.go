package model

type Todo struct {
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}
