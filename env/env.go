package env

import (
	"github.com/joho/godotenv"
	"log"
)

func InitializeDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file. Fallbacking to environmental variables.")
	} else {
		log.Println("Found .env file.")
	}
}
