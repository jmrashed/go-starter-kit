package logger

import (
	"go-starter-kit/pkg/config"
	"log"
)

func Init(config *config.Config) {
	// Initialize your logger here
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Logger initialized")
}
