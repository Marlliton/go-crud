package main

import (
	"github.com/Marlliton/gopportunities/config"
	"github.com/Marlliton/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
