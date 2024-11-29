package main

import (
	"go-rabbitmq/config"
	"go-rabbitmq/consumer"
	"go-rabbitmq/rabbitmq"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	config.LoadConfig()

	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 30 * time.Second,
	}))

	queueName := "consumer_queue"

	conn, ch, err := rabbitmq.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("RabbitMQ connection failed: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	if err := consumer.StartConsumer(ch, queueName); err != nil {
		log.Fatalf("Consumer not working: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "6060" // Default port
	}

	log.Fatal(app.Listen(":" + port))
}
