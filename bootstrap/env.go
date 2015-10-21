package bootstrap

import (
	"log"

	"gopkg.in/leyra/godotenv.v1"
)

func SetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
