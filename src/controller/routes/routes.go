package routes

import (
	"github.com/Marlliton/go-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	userGroup.GET("/:id", controller.FindUserByID)
	userGroup.GET("/email/:email", controller.FindUserByEmail)
	userGroup.POST("", controller.CreateUser)
	userGroup.PUT("", controller.UpdateUser)
	userGroup.DELETE("/:id", controller.DeleteUser)
}
