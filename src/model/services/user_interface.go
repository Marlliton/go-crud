package services

import (
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
	"github.com/Marlliton/go-crud/src/model/repository"
)

type userDomainService struct {
	repo repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr     // NOTE: Podemos omitir o nome dos  parametros, (string, userDomain) == (id string, user userDomain)
	FindUser(id string) (*model.UserDomainInterface, *rest_err.RestErr) // NOTE: Aqui já nomeamos o parametro
	DeleteUser(id string) *rest_err.RestErr
}

func NewUserDomainService(repo repository.UserRepository) UserDomainService {
	return &userDomainService{repo}
}
