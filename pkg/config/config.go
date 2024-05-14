package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	// Add the root directory to the config path
	viper.AddConfigPath(filepath.Join("..", ".."))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	config := &Config{
		ServerAddress: viper.GetString("server_address"),
	}

	return config
}
