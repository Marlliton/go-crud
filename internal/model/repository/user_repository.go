package repository

import (
	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/pkg/rest_err"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTION_NAME = "users"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(id string) *rest_err.RestErr
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
