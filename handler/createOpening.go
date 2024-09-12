package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		ctx.JSON(400, gin.H{"errors": err.Error()})
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errf("error creating opening %v", err.Error())
		return
	}
}
