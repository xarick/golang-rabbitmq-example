package pull

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"github.com/xarick/golang-rabbitmq-example/config"
)

func StartConsumer(cfg config.Application) error {
	conn, err := amqp.Dial(cfg.RMQUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		cfg.QueueName, // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received: %s", d.Body)
		}
	}()

	log.Println("Waiting for messages...")
	<-forever
	return nil
}
