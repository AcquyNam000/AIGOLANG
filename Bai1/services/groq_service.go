package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API Key viết trực tiếp (KHÔNG cần dùng .env nữa)
const apiKey = "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"

// ✅ Đúng URL API Groq
const groqAPIURL = "https://api.groq.com/openai/v1/chat/completions"

// Struct payload gửi API Groq
type groqPayload struct {
	Model    string         `json:"model"`
	Messages []MessageEntry `json:"messages"`
}

type MessageEntry struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Gọi API Groq
func CallGroqAPI(prompt string) (string, error) {
	if apiKey == "" {
		return "", errors.New("❌ API Key không hợp lệ hoặc trống")
	}

	// ✅ Sửa payload theo chuẩn Groq API
	payload := groqPayload{
		Model: "mixtral-8x7b-32768", // Chọn model phù hợp của Groq
		Messages: []MessageEntry{
			{"system", "Bạn là một trợ lý AI."},
			{"user", prompt},
		},
	}

	jsonData, _ := json.Marshal(payload)
	request, _ := http.NewRequest("POST", groqAPIURL, bytes.NewBuffer(jsonData))
	request.Header.Set("Authorization", "Bearer "+apiKey)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", errors.New("❌ Lỗi khi gọi API Groq")
	}
	defer resp.Body.Close()

	// ✅ Đọc dữ liệu trả về
	body, _ := ioutil.ReadAll(resp.Body)

	// ✅ In thử dữ liệu API Groq trả về để kiểm tra
	fmt.Println("🔍 API Response:", string(body))

	// ✅ Parse JSON để kiểm tra response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", errors.New("❌ Lỗi khi parse JSON từ API Groq")
	}

	// ✅ Kiểm tra response có dữ liệu hay không
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", errors.New("❌ API Groq không trả về dữ liệu hợp lệ")
	}

	firstChoice := choices[0].(map[string]interface{})
	message := firstChoice["message"].(map[string]interface{})
	content, exists := message["content"].(string)
	if !exists {
		return "", errors.New("❌ API Groq không trả về nội dung hợp lệ")
	}

	return content, nil
}
