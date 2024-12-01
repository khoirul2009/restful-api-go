package databases

import (
	"fmt"
	"learn-go-fiber/internal/config"
	"learn-go-fiber/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbConfig, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Could not load environment config: %v", err)
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Host, dbConfig.Port, dbConfig.SSLMode)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.SessionToken{})

	return db
}
