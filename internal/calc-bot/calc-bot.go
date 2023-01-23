package calcbot

import (
	"calc-bot/config"
	tgClient "calc-bot/internal/calc-bot/clients/telegram"
	event_consumer "calc-bot/internal/calc-bot/consumer/event-consumer"
	"calc-bot/internal/calc-bot/events/telegram"
	"calc-bot/internal/calc-bot/storage/sqlite"
	"context"
	"log"
)

const (
	batchSize   = 100
	storagePath = "./data/sqlite/storage.db"
)

func Start(config *config.Config) {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		log.Fatal("Service can't connect db: ", err)
	}

	if err := storage.Init(context.TODO()); err != nil {
		log.Fatal("Service can't init storage: ", err)
	}

	eventsHandler := telegram.New(tgClient.New(config.TG.TgHost, config.TG.TgToken), storage)

	log.Printf("Service has been started")

	consumer := event_consumer.New(eventsHandler, eventsHandler, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("Service stopped: ", err)
	}
}
