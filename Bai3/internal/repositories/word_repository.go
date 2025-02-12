package repositories

import (
	"database/sql"
	"Bai3/internal/models"
)

type WordRepository struct {
	DB *sql.DB
}

func NewWordRepository(db *sql.DB) *WordRepository {
	return &WordRepository{DB: db}
}

func (repo *WordRepository) SaveWord(word *models.Word) (int64, error) {
	query := "INSERT INTO word (lang, content, translate) VALUES ($1, $2, $3) RETURNING id"
	err := repo.DB.QueryRow(query, word.Lang, word.Content, word.Translate).Scan(&word.ID)
	return word.ID, err
}

func (repo *WordRepository) LinkWordToDialog(dialogID, wordID int64) error {
	query := "INSERT INTO word_dialog (dialog_id, word_id) VALUES ($1, $2)"
	_, err := repo.DB.Exec(query, dialogID, wordID)
	return err
}
