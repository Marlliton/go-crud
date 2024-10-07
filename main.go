package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Marlliton/go-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to loading .env file", err)
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("TESTE"))
}
