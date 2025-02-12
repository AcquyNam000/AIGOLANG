package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"net/http"
	"Bai3/config"
)
type GroqRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// ✅ Hàm gọi Groq API để tạo hội thoại
func GenerateDialog(prompt string) (string, error) {
	return sendRequest(prompt, "Tạo hội thoại")
}

func ExtractWords(dialog string) (string, error) {
	prompt := `Từ hội thoại sau, hãy trích xuất danh sách từ quan trọng cần học.
		Bỏ qua danh từ riêng. Trả về JSON đúng format: {"words": ["từ1", "từ2", "từ3"]}.
		Không giải thích gì thêm, chỉ xuất JSON.
		Hội thoại: ` + dialog

	response, err := sendRequest(prompt, "Trích xuất từ vựng")
	if err != nil {
		return "", err
	}

	// 🟢 Kiểm tra nếu response đã là JSON hợp lệ
	if json.Valid([]byte(response)) {
		fmt.Println("📌 JSON hợp lệ nhận được từ Groq:", response)
		return response, nil
	}

	// 🟢 Nếu không hợp lệ, lọc JSON bằng regex
	re := regexp.MustCompile(`\{.*\}`)
	matches := re.FindString(response)
	if matches == "" {
		fmt.Println("🚨 LỖI: Không tìm thấy JSON hợp lệ. Toàn bộ phản hồi:", response)
		return "", fmt.Errorf("Không tìm thấy JSON hợp lệ trong response")
	}

	fmt.Println("📌 JSON thực tế sau khi lọc:", matches)
	return matches, nil
}


func TranslateWords(words []string) (string, error) {
	wordJSON, _ := json.Marshal(map[string][]string{"words": words})
	prompt := `Dịch từng từ trong danh sách sau sang tiếng Anh.
		Trả về JSON đúng format {"translated_words": [{"vi": "từ tiếng Việt", "en": "từ tiếng Anh"}]}.
		Không giải thích gì thêm, chỉ xuất JSON.
		Danh sách từ: ` + string(wordJSON)

	response, err := sendRequest(prompt, "Dịch từ vựng")
	if err != nil {
		return "", err
	}

	// 🟢 Kiểm tra nếu response đã là JSON hợp lệ
	if json.Valid([]byte(response)) {
		fmt.Println("📌 JSON hợp lệ nhận được từ Groq:", response)
		return response, nil
	}

	// 🟢 Nếu không hợp lệ, lọc JSON bằng regex
	re := regexp.MustCompile(`\{.*\}`)
	matches := re.FindString(response)
	if matches == "" {
		fmt.Println("🚨 LỖI: Không tìm thấy JSON hợp lệ. Toàn bộ phản hồi:", response)
		return "", fmt.Errorf("Không tìm thấy JSON hợp lệ trong response")
	}

	fmt.Println("📌 JSON thực tế sau khi lọc:", matches)
	return matches, nil
}

// ✅ Hàm chung để gọi Groq API
func sendRequest(prompt, task string) (string, error) {
	cfg := config.GetConfig()

	requestBody, _ := json.Marshal(GroqRequest{
		Model: "llama3-8b-8192", // 🔹 Chọn model phù hợp (DeepSeek, Llama3, GPT-4o, ...)
		Messages: []Message{
			{Role: "system", Content: "Bạn là AI thực hiện tác vụ: " + task},
			{Role: "user", Content: prompt},
		},
	})

	req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+cfg.GroqAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to Groq API: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// ✅ Kiểm tra nếu API trả về lỗi
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Groq API error: %s - %s", resp.Status, string(body))
	}

	var groqResp GroqResponse
	if err := json.Unmarshal(body, &groqResp); err != nil {
		return "", fmt.Errorf("error parsing Groq API response: %v", err)
	}

	if len(groqResp.Choices) == 0 {
		return "", fmt.Errorf("Groq API returned no choices")
	}

	return groqResp.Choices[0].Message.Content, nil
}
