package main

import (
	"github.com/EugenSleptsov/triban/bot"
	"log"
)

func main() {
	botAPI, err := bot.NewBotAPI("YOUR_TELEGRAM_BOT_API_KEY")
	if err != nil {
		log.Panic(err)
	}

	botAPI.Start()
}
