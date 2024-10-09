package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Marlliton/go-crud/src/controller"
	"github.com/Marlliton/go-crud/src/controller/routes"
	"github.com/Marlliton/go-crud/src/model/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to loading .env file", err)
	}

	// NOTE: Init userController
	service := services.NewUserDomainService()
	userController := controller.NewUserController(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("TESTE"))
}
