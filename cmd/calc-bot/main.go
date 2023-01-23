package main

import (
	"context"
	"log"

	"calc-bot/config"
	tgClient "calc-bot/internal/clients/telegram"
	event_consumer "calc-bot/internal/consumer/event-consumer"
	"calc-bot/internal/events/telegram"
	"calc-bot/internal/storage/sqlite"
)

const (
	batchSize   = 100
	storagePath = "../../data/sqlite/storage.db"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("Can't get config: ", err)

	}

	storage, err := sqlite.New(storagePath)
	if err != nil {
		log.Fatal("Service can't connect db: ", err)
	}

	if err := storage.Init(context.TODO()); err != nil {
		log.Fatal("Service can't init storage: ", err)
	}

	eventsHandler := telegram.New(tgClient.New(config.TgHost, config.TgToken), storage)

	log.Printf("Service has been started")

	consumer := event_consumer.New(eventsHandler, eventsHandler, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("Service stopped: ", err)
	}
}
