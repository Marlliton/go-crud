package model

import (
	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
)

func (*userDomain) FindUser(string) (*userDomain, *rest_err.RestErr) {
	logger.Info("Init FindUser model", logger.Tag("journey", "FindUser"))

	return nil, nil
}
