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
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(string, model.UserDomainInterface) *rest_err.RestErr // NOTE: Podemos omitir o nome dos  parametros, (string, userDomain) == (id string, user userDomain)
	FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(id string) *rest_err.RestErr
}

func NewUserDomainService(repo repository.UserRepository) UserDomainService {
	return &userDomainService{repo}
}
