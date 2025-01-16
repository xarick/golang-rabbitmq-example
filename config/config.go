package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Application struct {
	RMQUrl    string
	QueueName string
}

func LoadConfig() Application {
	cfg := Application{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	cfg.RMQUrl = os.Getenv("RMQ_URL")
	cfg.QueueName = os.Getenv("QUEUE_NAME")

	return cfg
}
