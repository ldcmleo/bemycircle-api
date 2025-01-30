package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("[ERROR] loading .env file")
	}

	log.Println("[LOG] .env file loaded succesfully")
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		if defaultValue == "" {
			log.Fatalf("[ERROR] Variable %s is not loaded but needed", key)
		}

		return defaultValue
	}

	return value
}
