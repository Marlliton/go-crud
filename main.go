package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to loading .env file", err)
	}

	fmt.Println(os.Getenv("TESTE"))
}
