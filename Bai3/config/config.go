package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string
	DatabaseDSN string
	GroqAPIKey  string
}

var cfg Config

func Load() {
	viper.SetConfigName("config") // Chỉ cần tên file, không có .yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // Tìm trong thư mục config/
	viper.AddConfigPath(".")      // Tìm trong thư mục gốc

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	cfg = Config{
		ServerPort: viper.GetString("server.port"),
		DatabaseDSN: viper.GetString("database.dsn"),
		GroqAPIKey:  viper.GetString("groq.api_key"),
	}
}

func GetConfig() Config {
	return cfg
}
