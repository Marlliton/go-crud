package controller

import (
	"net/http"

	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/validation"
	"github.com/Marlliton/go-crud/src/controller/model/request"
	"github.com/Marlliton/go-crud/src/model"
	"github.com/Marlliton/go-crud/src/model/services"
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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Name,
		userRequest.Password,
		userRequest.Age,
	)

	service := services.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User created successfully",
		logger.Tag("journey", "CreateUser"),
	)

	ctx.String(http.StatusOK, "")
}
