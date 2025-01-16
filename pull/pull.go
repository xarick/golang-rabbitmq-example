package pull

import (
	"log"

	"github.com/xarick/golang-rabbitmq-example/config"
)

func StartPullService(cfg config.Application) error {
	log.Println("Pull Service Started")
	return StartConsumer(cfg)
}
