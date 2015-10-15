package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
)

func SetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
