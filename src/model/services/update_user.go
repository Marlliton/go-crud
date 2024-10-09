package services

import (
	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
)

func (*userDomainService) UpdateUserService(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateUser model", logger.Tag("journey", "UpdateUser"))

	return nil
}
