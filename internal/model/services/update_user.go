package services

import (
	"fmt"

	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/rest_err"
)

func (ud *userDomainService) UpdateUserService(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init UpdateUser model", logger.Tag("journey", "UpdateUser"))

	err := ud.repo.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error to call UpdateUserService", err, logger.Tag("journey", "UpdateUser"))
		return err
	}

	fmt.Println("Dentro do service mostrando o domain", userDomain)

	return nil
}
