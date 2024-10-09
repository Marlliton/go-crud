package services

import (
	"fmt"

	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
)

func (ud *userDomainService) CreateUser(
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
