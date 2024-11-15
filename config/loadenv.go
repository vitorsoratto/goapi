package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
