package routes

import (
	"github.com/Marlliton/go-crud/cmd/webserver/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, uc controller.UserControllerInterface) {
	userGroup := r.Group("/user")
	userGroup.GET("/:id", uc.FindUserByID)
	userGroup.GET("/email/:email", uc.FindUserByEmail)
	userGroup.POST("", uc.CreateUser)
	userGroup.PUT("/:id", uc.UpdateUser)
	userGroup.DELETE("/:id", uc.DeleteUser)
}
