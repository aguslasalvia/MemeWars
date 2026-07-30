package db

import (
	"log"
	"memewars/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("memewars.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	err = database.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Meme{},
		&models.Vote{},
	)

	if err != nil {
		log.Fatal("Failed running the migration: ", err)
	}

	DB = database
	log.Println("database connected && migration successfully")
}
