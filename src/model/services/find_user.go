package services

import (
	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
)

func (ud *userDomainService) FindUserByEmailService(email string) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init FindUserByEmail services", logger.Tag("journey", "FindUserByEmail"))

	return ud.repo.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByIDService(id string) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init FindUserByID services", logger.Tag("journey", "FindUserByID"))

	return ud.repo.FindUserByID(id)
}
