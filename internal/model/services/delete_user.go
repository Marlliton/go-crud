package services

import (
	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/rest_err"
)

func (ud *userDomainService) DeleteUserService(userId string) *rest_err.RestErr {

	logger.Info("Init DeleteUser model", logger.Tag("journey", "DeleteUser"))

	err := ud.repo.DeleteUser(userId)
	if err != nil {
		logger.Info("Init DeleteUser model", logger.Tag("journey", "DeleteUser"))
		return err
	}

	return nil
}
