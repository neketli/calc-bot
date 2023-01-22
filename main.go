package main

import (
	"log"
	"os"

	"calc-bot/clients/telegram"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	tgToken := os.Getenv("TG_TOKEN")
	tgHost := os.Getenv("TG_HOST")
	tgClient := telegram.New(tgHost, tgToken)

	//fetcher := fetcher.New(tgClient)
	//processor := processor.New(tgClient)
	//consumer.Start(fetcher. processor)
}
