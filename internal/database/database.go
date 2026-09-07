package database

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


var DB *gorm.DB

func InitialiseDB() {
    err := godotenv.Load()

    if err != nil {
        fmt.Println("Failed to load environment variables.")
        return
    }

    dsn := os.Getenv("DB_URL")

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        fmt.Println("Failed to connect to database.")
        return
    }

    fmt.Println("Database connected!")
}