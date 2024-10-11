package controller

import (
	"net/http"

	"github.com/Marlliton/go-crud/cmd/webserver/controller/dto/request"
	"github.com/Marlliton/go-crud/cmd/webserver/view"
	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/validation"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) CreateUser(ctx *gin.Context) {
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

	domainResult, err := uc.service.CreateUserService(domain)
	if err != nil {
		logger.Error(
			"Error trying to all CreateUser service", err,
			logger.Tag("journey", "CreateUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User created successfully",
		logger.Tag("journey", "CreateUser"),
	)

	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
