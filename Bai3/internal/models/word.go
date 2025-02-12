package models

type Word struct {
	ID        int64  `json:"id"`
	Lang      string `json:"lang"`
	Content   string `json:"vi"`  // 🔹 Đảm bảo khớp với key JSON từ API Groq
	Translate string `json:"en"`  // 🔹 Đảm bảo khớp với key JSON từ API Groq
}
