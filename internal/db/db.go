package db

import (
	"fmt"
	"log"
	"os"

	"github.com/DigdarshanMohanty/notes-api/internal/models"
	"gorm.io/driver/postgres"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSLMODE"),
    )
	fmt.Println("DSN:", dsn)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }
    err = db.AutoMigrate(&models.User{}, &models.Note{})
    if err != nil {
        log.Fatal("Failed to migrate DB:", err)
    }
    return db
}
