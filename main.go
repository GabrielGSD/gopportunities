package main

import (
	"github.com/GabrielGSD/gopportunities/config"
	"github.com/GabrielGSD/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize the configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		// panic(err)
		return
	}

	router.Initialize()
}
