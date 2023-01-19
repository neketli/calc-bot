package main

import (
	"log"

	"github.com/joho/godotenv"
	//"os"
)

func main() {
	//var tgClient Client

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	//tgToken := os.Getenv("TG_TOKEN")
	//tgClient = telegram.New(tgToken)

	//fetcher := fetcher.New(tgClient)
	//processor := processor.New(tgClient)
	//consumer.Start(fetcher. processor)
}
