package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"carrotfarmer/chad-stack/models/todo"
	"carrotfarmer/chad-stack/models/user"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("deez_nuts.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&user.User{}, &todo.Todo{})
	if err != nil {
		log.Fatalf("ERROR: database migration failed: %v", err)
	}

	DB = database
}
