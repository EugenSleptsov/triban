package main

import (
	"github.com/EugenSleptsov/triban/bot"
	"log"
)

func main() {
	botAPI, err := bot.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	botAPI.Start()
}
