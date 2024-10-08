package model

import (
	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
)

func (*userDomain) UpdateUser(string) *rest_err.RestErr {
	logger.Info("Init UpdateUser model", logger.Tag("journey", "UpdateUser"))

	return nil
}
