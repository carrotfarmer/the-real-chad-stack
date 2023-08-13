package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"carrotfarmer/chad-stack/auth"
	"carrotfarmer/chad-stack/models/todo"
	"carrotfarmer/chad-stack/models/user"
	"carrotfarmer/chad-stack/router"
)

func initDB() {
	db, err := gorm.Open(sqlite.Open("deez_nuts.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	db.AutoMigrate(&user.User{}, &todo.Todo{})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load the env vars: %v", err)
	}

	auth, err := auth.New()
	if err != nil {
		log.Fatalf("failed to initialize the authenticator: %v", err)
	}

	r := router.New(auth)

	log.Print("server up on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("ERROR: server crashed: %v", err)
	}
}
