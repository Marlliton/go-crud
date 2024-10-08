package services

import (
	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
)

func (*userDomainService) DeleteUser(string) *rest_err.RestErr {
	logger.Info("Init deleteUser model", logger.Tag("journey", "deleteUser"))

	return nil
}
