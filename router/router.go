package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
  // Create a new router
	r := gin.Default()

  // Initialize routes
	initRoutes(r)

  // Start the server
	r.Run(":8080")
}
