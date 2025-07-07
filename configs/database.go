package configs

import (
	"os"
	"estadisticas-api/internal/shared/database"
)

func GetDatabaseConfig() database.Config {
	return database.Config{
		Host:     getEnv("DB_HOST",""),
		Port:     getEnv("DB_PORT", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		Database: getEnv("DB_NAME",""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}