package database

import (
	"log"
	"os"

	"urlshortener/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func Connect() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "urlshortener.db"
	}

	// Use modernc.org/sqlite for pure Go SQLite driver
	database, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        dbPath,
	}, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.URL{})
	DB = database
	log.Println("Database connected successfully")
}
