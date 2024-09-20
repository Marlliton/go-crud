package handler

import (
	"net/http"

	"github.com/Marlliton/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "validation error")
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Link:     request.Link,
		Remote:   *request.Remote,
		Company:  request.Company,
		Salary:   request.Salary,
		Location: request.Location,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating opening -> %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on databse")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
