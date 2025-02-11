package models

// GroqRequest chứa dữ liệu request từ client
type GroqRequest struct {
	Prompt string `json:"prompt"`
}
