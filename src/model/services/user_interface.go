package services

import (
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
)

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr     // NOTE: Podemos omitir o nome dos  parametros, (string, userDomain) == (id string, user userDomain)
	FindUser(id string) (*model.UserDomainInterface, *rest_err.RestErr) // NOTE: Aqui já nomeamos o parametro
	DeleteUser(id string) *rest_err.RestErr
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
