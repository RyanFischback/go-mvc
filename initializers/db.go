package initializers

import (
	"fmt"
	"mvc/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
	} else {
		fmt.Println("Successfully connected to database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Post{}) // Create the Table
}
