package main

import (
	"log"

	"github.com/alfaluizh/golang-rabbitmq-producer/internal/http"
	"github.com/alfaluizh/golang-rabbitmq-producer/internal/rabbitmq"
	"github.com/alfaluizh/golang-rabbitmq-producer/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	producer := rabbitmq.NewProducer(cfg.RabbitmqURL, cfg.RabbitmqQueue, cfg.RabbitmqReplyQueue)

	router := gin.Default()
	handler := http.NewHandler(producer)
	router.POST("/publish", handler.Publish)
	log.Println("Server Running on :8080")

	router.Run(":8080")
}
