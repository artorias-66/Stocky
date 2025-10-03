package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"stock-reward-system/src/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db
	DB.AutoMigrate(&models.Reward{}, &models.LedgerEntry{}, &models.StockPrice{})
}