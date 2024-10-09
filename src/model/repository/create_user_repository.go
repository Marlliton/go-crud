package repository

import (
	"context"

	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
)

const (
	COLLECTION_NAME = "users"
)

func (ur *useRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	collection := ur.databaseConnection.Collection(COLLECTION_NAME)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerErr(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerErr(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
