package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	JWTSecret string
	AppPort   string
}

func Load() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		DBHost:    get("DB_HOST", os.Getenv("DB_HOST")),
		DBPort:    get("DB_PORT", os.Getenv("DB_PORT")),
		DBUser:    get("DB_USER", os.Getenv("DB_USER")),
		DBPass:    get("DB_PASSWORD", os.Getenv("DB_PASSWORD")),
		DBName:    get("DB_NAME", os.Getenv("DB_NAME")),
		JWTSecret: get("JWT_SECRET", os.Getenv("JWT_SECRET")),
		AppPort:   get("APP_PORT", os.Getenv("PORT")),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}
	return cfg
}

func get(k, fallback string) string {
	v := os.Getenv(k)
	if v == "" {
		return fallback
	}
	return v
}
