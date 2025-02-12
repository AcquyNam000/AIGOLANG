package models

type Dialog struct {
	ID      int64  `json:"id"`
	Lang    string `json:"lang"`
	Content string `json:"content"`
}
