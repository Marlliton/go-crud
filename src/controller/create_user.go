package controller

import (
	"net/http"

	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/validation"
	"github.com/Marlliton/go-crud/src/controller/model/request"
	"github.com/Marlliton/go-crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	logger.Info(
		"Init CreateUser controller",
		logger.Tag("journey", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validate user info", err,
			logger.Tag("journey", "CreateUser"),
		)
		restErr := validation.ValidateuserErr(err)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "teste",
		Name:  userRequest.Name,
		Age:   userRequest.Age,
		Email: userRequest.Email,
	}

	logger.Info(
		"User created successfully",
		logger.Tag("journey", "CreateUser"),
	)

	ctx.JSON(http.StatusOK, response)
}
