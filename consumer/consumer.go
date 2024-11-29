package consumer

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func StartConsumer(ch *amqp091.Channel, queueName string) error {
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			log.Printf("Recieved message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Listen queue message: %s. For exit CTRL+C", queueName)
	select {}
}
