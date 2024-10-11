package controller

import (
	"net/http"

	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) DeleteUser(ctx *gin.Context) {

	logger.Info(
		"Init DeleteUser controller",
		logger.Tag("journey", "DeleteUser"),
	)

	userId := ctx.Param("id")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestErr("Invalid userId.")
		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error(
			"Error trying to all DeleteUser service", err,
			logger.Tag("journey", "DeleteUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User created successfully",
		logger.Tag("journey", "DeleteUser"),
	)

	ctx.Status(http.StatusOK)
}
