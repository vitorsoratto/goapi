package main

import (
	"fmt"

	"goapi/config"
	"goapi/router"
)

func main() {
	// Connect to the database
	err := config.ConnectDB()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
	}

	// Initialize the router
	router.InitRouter()
}
