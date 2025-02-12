package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"Bai3/internal/models"
	"Bai3/internal/repositories"
	"Bai3/pkg/groq"
)

type DialogService struct {
	DialogRepo *repositories.DialogRepository
	WordRepo   *repositories.WordRepository
}

func NewDialogService(dialogRepo *repositories.DialogRepository, wordRepo *repositories.WordRepository) *DialogService {
	return &DialogService{DialogRepo: dialogRepo, WordRepo: wordRepo}
}

func (s *DialogService) ProcessDialog(prompt string) (*models.Dialog, []models.Word, error) {
	// 🟢 Gọi API tạo hội thoại
	dialogContent, err := groq.GenerateDialog(prompt)
	if err != nil {
		return nil, nil, err
	}

	// 🟢 Lưu hội thoại vào database
	dialog := &models.Dialog{Lang: "vi", Content: dialogContent}
	dialogID, err := s.DialogRepo.SaveDialog(dialog)
	if err != nil {
		return nil, nil, err
	}
	dialog.ID = dialogID

	// 🟢 Gọi API trích xuất từ
	wordsJSON, err := groq.ExtractWords(dialogContent)
	if err != nil {
		return nil, nil, err
	}

	// 🟢 Parse JSON trích xuất từ
	var wordList struct {
		Words []string `json:"words"`
	}
	if err := json.Unmarshal([]byte(wordsJSON), &wordList); err != nil {
		return nil, nil, errors.New("invalid word extraction response")
	}

	// 🟢 Gọi API dịch từ
	translatedWordsJSON, err := groq.TranslateWords(wordList.Words)
	if err != nil {
		return nil, nil, err
	}

	// 🟢 Parse JSON dịch từ
	var translatedWords struct {
		TranslatedWords []models.Word `json:"translated_words"`
	}
	if err := json.Unmarshal([]byte(translatedWordsJSON), &translatedWords); err != nil {
		fmt.Println("🚨 LỖI Parse JSON dịch từ:", err)
		fmt.Println("📌 API trả về:", translatedWordsJSON)
		return nil, nil, errors.New("invalid translation response")
	}

	// ✅ Log chi tiết trước khi lưu từ
	fmt.Println("📌 Danh sách từ sau khi parse:", translatedWords.TranslatedWords)

	// 🟢 Lưu từ vào database và liên kết với hội thoại
	savedWords := []models.Word{}
	for _, word := range translatedWords.TranslatedWords {
		// ✅ Kiểm tra dữ liệu trước khi lưu
		fmt.Println("📌 Từ đang xử lý:", word)

		if word.Content == "" || word.Translate == "" {
			fmt.Println("🚨 CẢNH BÁO: Từ bị rỗng, bỏ qua:", word)
			continue
		}

		// ✅ Lưu từ vào database
		word.Lang = "vi"
		wordID, err := s.WordRepo.SaveWord(&word)
		if err != nil {
			fmt.Println("🚨 LỖI khi lưu từ:", word, err)
			return nil, nil, err
		}

		// ✅ Liên kết từ với hội thoại
		err = s.WordRepo.LinkWordToDialog(dialogID, wordID)
		if err != nil {
			return nil, nil, err
		}

		// ✅ Thêm từ vào danh sách trả về
		word.ID = wordID
		savedWords = append(savedWords, word)
	}

	fmt.Println("📌 Hoàn tất xử lý hội thoại ID:", dialogID)
	return dialog, savedWords, nil
}
