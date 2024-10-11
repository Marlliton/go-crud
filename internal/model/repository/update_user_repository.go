package repository

import (
	"context"

	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/internal/model/repository/entity/converter"
	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *useRepository) UpdateUser(
	userId string, userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateUser repository", logger.Tag("journey", "UpdateUser"))

	collection := ur.databaseConnection.Collection(COLLECTION_NAME)
	value := converter.ConvetDomainToEntity(userDomain)

	objID, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$set", Value: value}}
	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		logger.Info("Error trying to UpdateUser", logger.Tag("journey", "UpdateUser"))

		return rest_err.NewInternalServerErr(err.Error())
	}

	return nil
}
