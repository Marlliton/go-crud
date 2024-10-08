package model

import (
	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
)

func (*userDomain) DeleteUser(string) *rest_err.RestErr {
	logger.Info("Init deleteUser model", logger.Tag("journey", "deleteUser"))

	return nil
}
