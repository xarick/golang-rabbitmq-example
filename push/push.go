package push

import (
	"log"

	"github.com/xarick/golang-rabbitmq-example/config"
)

func StartPushService(cfg config.Application) error {
	log.Println("Push service started")
	return StartProducer(cfg)
}
