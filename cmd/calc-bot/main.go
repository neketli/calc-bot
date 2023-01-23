package main

import (
	"log"

	"calc-bot/config"
	calcbot "calc-bot/internal/calc-bot"
)

const configPath = "./config/config.yml"

func main() {
	config, err := config.New(configPath)
	if err != nil {
		log.Fatal("Can't setup config: ", err)
	}
	calcbot.Start(config)
}
