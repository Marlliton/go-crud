package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Marlliton/go-crud/cmd/webserver/controller"
	"github.com/Marlliton/go-crud/cmd/webserver/controller/routes"
	"github.com/Marlliton/go-crud/internal/model/repository"
	"github.com/Marlliton/go-crud/internal/model/services"
	"github.com/Marlliton/go-crud/pkg/database/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to loading .env file", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	database, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	// NOTE: Init userController
	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("TESTE"))
}

func initDependencies(database *mongo.Database) controller.UserControllerInterface {

	repo := repository.NewUserRepository(database)
	service := services.NewUserDomainService(repo)
	userController := controller.NewUserController(service)

	return userController
}
