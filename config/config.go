package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configuration ...
// type Configuration struct {
// 	Application struct {
// 		Port string
// 	}

// 	Database struct {

// 	}
// }

// var cfg *Configuration

// Get ...
func Get(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

// Get environment variable
// func Get() (c *Configuration) {
// 	return cfg
// }
