package main

import (
	"github.com/EugenSleptsov/triban/bot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the TELEGRAM_TOKEN from the environment
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		log.Fatalf("TELEGRAM_TOKEN is not set in the environment")
	}

	botAPI, err := bot.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Start()
}
