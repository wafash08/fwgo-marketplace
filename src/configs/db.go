package configs

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	url := os.Getenv("DB_URL")
	dialect := postgres.Open(url)
	var err error
	DB, err = gorm.Open(dialect, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}
