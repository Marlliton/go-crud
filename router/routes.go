package router

import (
	"github.com/Marlliton/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	rGroup := r.Group("/api/v1")
	{
		rGroup.GET("/opening", handler.ShowOpeningHandler)
		rGroup.GET("/openings", handler.ListOpeningsHandler)
		rGroup.POST("/opening", handler.CreateOpeningHandler)
		rGroup.DELETE("/opening", handler.DeleteOpeningHandler)
		rGroup.PUT("/opening", handler.UpdateOpeningHandler)
	}
}
