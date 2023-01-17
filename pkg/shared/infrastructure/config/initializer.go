package config

import (
	"github.com/joho/godotenv"
	"log"
)

// InitializeDotEnv Load ENV file
func InitializeDotEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
