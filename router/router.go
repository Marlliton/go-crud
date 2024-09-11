package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	// Initialize routes
	initializeRoutes(r)
	r.Run(":8080")
}
