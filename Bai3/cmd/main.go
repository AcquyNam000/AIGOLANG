package main

import (
	"database/sql"
	"log"
	"Bai3/config"
	"Bai3/internal/handlers"
	"Bai3/internal/repositories"
	"Bai3/internal/router"
	"Bai3/internal/services"

	_ "github.com/lib/pq"
)

func main() {
	config.Load()

	db, err := sql.Open("postgres", config.GetConfig().DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dialogRepo := repositories.NewDialogRepository(db)
	wordRepo := repositories.NewWordRepository(db)
	dialogService := services.NewDialogService(dialogRepo, wordRepo)
	dialogHandler := handlers.NewDialogHandler(dialogService)

	r := router.SetupRouter(dialogHandler)
	r.Run(":8080")
}
