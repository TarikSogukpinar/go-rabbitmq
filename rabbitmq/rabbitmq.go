package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ() (*amqp091.Connection, *amqp091.Channel, error) {
	url := os.Getenv("RABBITMQ_URL")

	conn, err := amqp091.DialConfig(url, amqp091.Config{
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	log.Println("Connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return conn, ch, nil
}
