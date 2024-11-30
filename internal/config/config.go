// internal/config/config.go
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
	SSLMode  string
}

// LoadEnv loads environment variables from .env file
func LoadEnv() (*DatabaseConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	config := &DatabaseConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	return config, nil
}
