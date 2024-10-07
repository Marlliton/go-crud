package controller

import (
	"net/http"

	"github.com/Marlliton/go-crud/src/configuration/validation"
	"github.com/Marlliton/go-crud/src/controller/model/request"
	"github.com/Marlliton/go-crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
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

	ctx.JSON(http.StatusOK, response)
}
