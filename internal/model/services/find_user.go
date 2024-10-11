package services

import (
	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/rest_err"
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
