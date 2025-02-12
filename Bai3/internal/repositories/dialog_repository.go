package repositories

import (
	"database/sql"
	"Bai3/internal/models"
)

type DialogRepository struct {
	DB *sql.DB
}

func NewDialogRepository(db *sql.DB) *DialogRepository {
	return &DialogRepository{DB: db}
}

func (repo *DialogRepository) SaveDialog(dialog *models.Dialog) (int64, error) {
	query := "INSERT INTO dialog (lang, content) VALUES ($1, $2) RETURNING id"
	err := repo.DB.QueryRow(query, dialog.Lang, dialog.Content).Scan(&dialog.ID)
	return dialog.ID, err
}
