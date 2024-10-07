package main

import "github.com/Marlliton/go-crud/config"

var (
	logger *config.Logger
)

func main() {
	// Sample logger
	logger = config.GetLogger("main")
}
