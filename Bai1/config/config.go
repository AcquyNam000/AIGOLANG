package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load biến môi trường từ file .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
	}
}

// Hàm lấy biến môi trường với giá trị mặc định
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
