package services

import (
	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
)

func (*userDomainService) FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUser model", logger.Tag("journey", "FindUser"))

	return nil, nil
}
