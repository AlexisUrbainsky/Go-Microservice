package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadVariables() {

	if os.Getenv("ENV") != "PROD" {

		err := godotenv.Load("../.env")

		if err != nil {
			log.Fatal("error reading env file")
		}

	}

}
