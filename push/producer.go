package push

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
	"github.com/xarick/golang-rabbitmq-example/config"
)

func StartProducer(cfg config.Application) error {
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

	q, err := ch.QueueDeclare(
		cfg.QueueName, // queue name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := 0; i < 120; i++ {
		num := r.Intn(1000)
		body := fmt.Sprintf("%d", num)

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			return fmt.Errorf("failed to publish a message: %v", err)
		}
		log.Printf("Sent: %s", body)
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
