package config

import (
	"fmt"
	"log"
	"mediportal/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️  .env file not found, skipping... assuming env vars are preloaded")
	}

	dbURL := os.Getenv("DB_URL")
	fmt.Println("DB_URL:", os.Getenv("DB_URL"))

	if dbURL == "" {
		log.Fatal("DB_URL not set in environment variables")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	DB = db

	var dbName string
	err = db.Raw("SELECT current_database()").Scan(&dbName).Error
	if err != nil {
		log.Fatal("Failed to retrieve database name:", err)
	}
	fmt.Println("Connected to DB:", dbName)

	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		log.Fatal("Error enabling UUID extension:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Patient{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Connected to DB and migrated!")
}
