package services

import (
	"fmt"

	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/rest_err"
)

func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model", logger.Tag("journey", "createUser"))

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.repo.CreateUser(userDomain)
	if err != nil {
		logger.Info("Init createUser model", logger.Tag("journey", "CreateUser"))
		return nil, err
	}

	fmt.Println("Dentro do service mostrando o domain", userDomain)

	return userDomainRepository, nil
}
