package main

import (
	"log"
	"time"

	"github.com/xarick/golang-rabbitmq-example/config"
	"github.com/xarick/golang-rabbitmq-example/pull"
	"github.com/xarick/golang-rabbitmq-example/push"
)

func main() {
	cfg := config.LoadConfig()

	go func() {
		if err := push.StartPushService(cfg); err != nil {
			log.Fatalf("Push service failed: %v", err)
		}
	}()

	// queue qo'shib olishi uchun ozgina kutamiz
	time.Sleep(100 * time.Millisecond)

	go func() {
		if err := pull.StartPullService(cfg); err != nil {
			log.Fatalf("Pull service failed: %v", err)
		}
	}()

	// Servislarni bloklaymiz
	select {}
}
