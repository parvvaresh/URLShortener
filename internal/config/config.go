package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl   string
	Port    string
	BaseURL string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		DBUrl:   getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/urlshortener?sslmode=disable"),
		Port:    getEnv("PORT", "8080"),
		BaseURL: getEnv("BASE_URL", "http://localhost:8080"),
	}

	log.Println("Config loaded:", cfg)
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
