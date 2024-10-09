package repository

import (
	"context"
	"fmt"

	"github.com/Marlliton/go-crud/src/configuration/logger"
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
	"github.com/Marlliton/go-crud/src/model/repository/entity"
	"github.com/Marlliton/go-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *useRepository) FindUserByEmail(email string) (
	model.UserDomainInterface, *rest_err.RestErr,
) {

	logger.Info("Init FindUserByEmail repository", logger.Tag("journey", "FindUserByEmail"))

	collection := ur.databaseConnection.Collection(COLLECTION_NAME)

	userEntity := &entity.UserEntity{}

	// NOTE: Crinado filtro para consulta no mongoDB
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found to this email: %s", email)
			logger.Error(
				errorMessage,
				err,
				logger.Tag("journey", "findUser"),
			)

			return nil, rest_err.NewNotFoundErr(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(
			errorMessage,
			err,
			logger.Tag("journey", "findUser"),
		)

		return nil, rest_err.NewInternalServerErr(errorMessage)
	}

	logger.Info(
		"FindUser by email successfully completed",
		logger.Tag("journey", "findUser"),
		logger.Tag("email", email),
	)

	return converter.ConvetEntityToDomain(*userEntity), nil

}

func (ur *useRepository) FindUserByID(id string) (
	model.UserDomainInterface, *rest_err.RestErr,
) {

	logger.Info("Init FindUserByID repository", logger.Tag("journey", "FindUserByID"))

	collection := ur.databaseConnection.Collection(COLLECTION_NAME)

	userEntity := &entity.UserEntity{}

	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objID}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found to this ID: %s", id)
			logger.Error(
				errorMessage,
				err,
				logger.Tag("journey", "findUser"),
			)

			return nil, rest_err.NewNotFoundErr(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		logger.Error(
			errorMessage,
			err,
			logger.Tag("journey", "findUser"),
		)

		return nil, rest_err.NewInternalServerErr(errorMessage)
	}

	logger.Info(
		"FindUser by ID successfully completed",
		logger.Tag("journey", "findUser"),
		logger.Tag("id", userEntity.ID),
	)

	return converter.ConvetEntityToDomain(*userEntity), nil

}
