package main

import (
	"go-starter-kit/pkg/config"
	"go-starter-kit/pkg/logger"
	"go-starter-kit/pkg/server"
)

func main() {
	// Initialize configuration
	config := config.LoadConfig()

	// Initialize logger
	logger.Init(config)

	// Initialize server
	r := server.NewRouter()

	// Start server
	r.Run(config.ServerAddress)
}
