package main

import (
	"log"

	"github.com/nlnaa11/FinanceTelegramBot/internal/clients/tg"
	"github.com/nlnaa11/FinanceTelegramBot/internal/config"
	"github.com/nlnaa11/FinanceTelegramBot/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed: ", err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}
