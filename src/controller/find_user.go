package controller

import (
	"net/http"
	"net/mail"

	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) FindUserByID(ctx *gin.Context) {
	logger.Info(
		"Init FindUserByID controller",
		logger.Tag("journey", "FindUser"),
	)
	id := ctx.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		errorMessage := rest_err.NewBadRequestErr("Invalid user id.")

		logger.Error(
			"Error on validate userID",
			err,
			logger.Tag("journey", "FindUser"),
		)
		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	domain, err := uc.service.FindUserByIDService(id)
	if err != nil {

		logger.Error(
			"Error on call FindaUserByID service",
			err,
			logger.Tag("journey", "FindUserByID"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}

func (uc *userControllerInterface) FindUserByEmail(ctx *gin.Context) {
	logger.Info(
		"Init FindUserByEmail controller",
		logger.Tag("journey", "FindUser"),
	)
	email := ctx.Param("email")
	if _, err := mail.ParseAddress(email); err != nil {
		errorMessage := rest_err.NewBadRequestErr("Invalid user email.")

		logger.Error(
			"Error on validate userEmail",
			err,
			logger.Tag("journey", "FindUserByEmail"),
		)
		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	domain, err := uc.service.FindUserByEmailService(email)
	if err != nil {

		logger.Error(
			"Error on call FindaUserByID service",
			err,
			logger.Tag("journey", "FindUserByEmail"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
