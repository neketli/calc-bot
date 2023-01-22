package main

import (
	"log"
	"os"

	tgClient "calc-bot/clients/telegram"
	event_consumer "calc-bot/consumer/event-consumer"
	"calc-bot/events/telegram"

	"github.com/joho/godotenv"
)

const (
	batchSize = 100
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	tgToken := os.Getenv("TG_TOKEN")
	tgHost := os.Getenv("TG_HOST")

	eventsHandler := telegram.New(tgClient.New(tgHost, tgToken))

	log.Printf("Service has been started")

	consumer := event_consumer.New(eventsHandler, eventsHandler, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("Service stopped: ", err)
	}
}
