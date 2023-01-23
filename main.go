package main

import (
	"context"
	"log"
	"os"

	tgClient "calc-bot/clients/telegram"
	event_consumer "calc-bot/consumer/event-consumer"
	"calc-bot/events/telegram"
	"calc-bot/storage/sqlite"

	"github.com/joho/godotenv"
)

const (
	batchSize   = 100
	storagePath = "data/sqlite"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	tgToken := os.Getenv("TG_TOKEN")
	tgHost := os.Getenv("TG_HOST")

	storage, err := sqlite.New(storagePath)
	if err != nil {
		log.Fatal("Service can't connect db: ", err)
	}

	if err := storage.Init(context.TODO()); err != nil {
		log.Fatal("Service can't init storage: ", err)
	}

	eventsHandler := telegram.New(tgClient.New(tgHost, tgToken), storage)

	log.Printf("Service has been started")

	consumer := event_consumer.New(eventsHandler, eventsHandler, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("Service stopped: ", err)
	}
}
