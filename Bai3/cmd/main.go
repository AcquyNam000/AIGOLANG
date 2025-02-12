package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"Bai3/config"
	"Bai3/internal/handlers"
	"Bai3/internal/repositories"
	"Bai3/internal/services"

	_ "github.com/lib/pq"
)

func main() {
	// Load config
	config.Load()

	// Kết nối database
	db, err := sql.Open("postgres", config.GetConfig().DatabaseDSN)
	if err != nil {
		log.Fatal("Lỗi kết nối database:", err)
	}
	defer db.Close()

	// Khởi tạo repository & service
	dialogRepo := repositories.NewDialogRepository(db)
	wordRepo := repositories.NewWordRepository(db)
	dialogService := services.NewDialogService(dialogRepo, wordRepo)
	dialogHandler := handlers.NewDialogHandler(dialogService)

	// 🔥 Tạo router GIN
	r := gin.Default()

	// 🔥 Thêm CORS Middleware trước khi đăng ký route
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Chỉ cho phép ReactJS
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 🔥 Đăng ký route
	api := r.Group("/api/dialog")
	{
		api.POST("/process", dialogHandler.ProcessDialog)
		api.POST("/manual", dialogHandler.ProcessManualDialog)
	}

	// 🔥 Chạy server
	port := ":8080"
	log.Println("✅ Server đang chạy tại", port)
	r.Run(port)
}
