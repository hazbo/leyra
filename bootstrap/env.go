package bootstrap

import (
	"gopkg.in/leyra/godotenv.v1"
	"log"
)

func SetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
