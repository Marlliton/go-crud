package repository

import (
	"github.com/Marlliton/go-crud/src/configuration/rest_err"
	"github.com/Marlliton/go-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTION_NAME = "users"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
}

type useRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &useRepository{
		databaseConnection: database,
	}
}
