package main

import (
	"fmt"

	"goapi/config"
	"goapi/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	// Connect to the database
	err = config.ConnectDB()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
	}

	// Initialize the router
	router.InitRouter()
}
