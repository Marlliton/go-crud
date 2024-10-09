package main

import (
	"github.com/Marlliton/go-crud/src/controller"
	"github.com/Marlliton/go-crud/src/model/repository"
	"github.com/Marlliton/go-crud/src/model/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {

	repo := repository.NewUserRepository(database)
	service := services.NewUserDomainService(repo)
	userController := controller.NewUserController(service)

	return userController
}
