package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// InitializeDotEnv Load ENV file
func InitializeDotEnv() {
	execPath, err := os.Executable()

	if err != nil {
		log.Fatal("Error setting execPath > " + err.Error())
	}

	err = godotenv.Load(filepath.Dir(execPath) + "/../.env")

	if err != nil {
		log.Fatal("Error loading .env file > " + err.Error())
	}
}
