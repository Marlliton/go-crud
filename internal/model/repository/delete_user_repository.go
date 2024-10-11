package repository

import (
	"context"

	"github.com/Marlliton/go-crud/pkg/logger"
	"github.com/Marlliton/go-crud/pkg/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *useRepository) DeleteUser(userId string) *rest_err.RestErr {

	logger.Info("Init DeleteUser repository", logger.Tag("journey", "DeleteUser"))

	collection := ur.databaseConnection.Collection(COLLECTION_NAME)

	objID, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: objID}}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		logger.Info("Error trying to DeleteUser", logger.Tag("journey", "DeleteUser"))

		return rest_err.NewInternalServerErr(err.Error())
	}

	return nil
}
