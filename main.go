package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Marlliton/go-crud/src/configuration/database/mongodb"
	"github.com/Marlliton/go-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to loading .env file", err)
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
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
