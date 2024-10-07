package controller

import (
	"fmt"

	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestErr(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()),
		)

		ctx.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
