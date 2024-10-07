package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	userGroup.GET("/:id")
	userGroup.GET("/:email")
	userGroup.POST("/")
	userGroup.PUT("/")
	userGroup.DELETE("/:id")
}
