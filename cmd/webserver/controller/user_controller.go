package controller

import (
	"github.com/Marlliton/go-crud/internal/model/services"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	CreateUser(*gin.Context)
	DeleteUser(*gin.Context)
	FindUserByID(*gin.Context)
	FindUserByEmail(*gin.Context)
	UpdateUser(*gin.Context)
}

type userControllerInterface struct {
	service services.UserDomainService
}

func NewUserController(service services.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service,
	}
}
