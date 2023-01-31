package config

import (
	"github.com/joho/godotenv"
	"log"
	"github.com/sirupsen/logrus"
)

func InitEnvs(filepath string) {
	err := godotenv.Load(filepath)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logrus.Infof("File .env successfully loaded\n")
}
