package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system envs")
	}

	// Build DSN
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("DB connection failed: %v", err)
	}
	DB = db

	fmt.Println("Connected to DB: ", os.Getenv("DB_NAME"))

	return db, nil
	// Ping to test connection
	// if err := db.Ping(); err != nil {
	// 	return nil, fmt.Errorf("error connection: %v", err)
	// }
	// fmt.Println("Connected to DB: ", os.Getenv("DB_NAME"))
	// return &gorm.DB{}, nil
	// }
}
