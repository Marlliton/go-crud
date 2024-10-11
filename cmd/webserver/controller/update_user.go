package controller

import (
	"net/http"
	"strings"

	"github.com/Marlliton/go-crud/cmd/webserver/controller/dto/request"
	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/validation"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) UpdateUser(ctx *gin.Context) {

	logger.Info(
		"Init UpdateUser controller",
		logger.Tag("journey", "UpdateUser"),
	)
	var userUpdateRequest request.UserUpdateRequest
	userId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error(
			"Error trying to validate user info", err,
			logger.Tag("journey", "UpdateUser"),
		)
		restErr := validation.ValidateuserErr(err)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userUpdateRequest.Name,
		userUpdateRequest.Age,
	)

	err := uc.service.UpdateUserService(userId, domain)
	if err != nil {
		logger.Error(
			"Error trying to all UpdateUser service", err,
			logger.Tag("journey", "UpdateUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User created successfully",
		logger.Tag("journey", "UpdateUser"),
	)

	ctx.Status(http.StatusOK)
}
