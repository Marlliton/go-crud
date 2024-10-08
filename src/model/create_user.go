package model

import (
	"fmt"

	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
)

func (u *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", logger.Tag("journey", "createUser"))

	u.EncryptPassword()

	fmt.Println(u)

	return nil
}
