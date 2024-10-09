package repository

import (
	"context"

	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
	"github.com/Marlliton/go-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *useRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", logger.Tag("journey", "createUser"))

	collection := ur.databaseConnection.Collection(COLLECTION_NAME)

	value := converter.ConvetDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerErr(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return converter.ConvetEntityToDomain(*value), nil
}
