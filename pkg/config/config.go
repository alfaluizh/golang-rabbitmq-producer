package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitmqURL   string
	RabbitmqQueue string
	RabbitmqReplyQueue string
}

func Load() Config {
	godotenv.Load()

	return Config{
		RabbitmqURL:        os.Getenv("RABBITMQ_URL"),
		RabbitmqQueue:      os.Getenv("RABBITMQ_QUEUE"),
		RabbitmqReplyQueue: os.Getenv("RABBITMQ_REPLY_QUEUE"),
	}
}
