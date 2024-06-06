package bot

import (
	"fmt"
	"github.com/EugenSleptsov/triban/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	commands map[string]commands.Command
}

func NewBotAPI(apiKey string) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		api: botAPI,
		commands: map[string]commands.Command{
			"/help":   commands.HelpCommand{},
			"/ininal": commands.IninalCommand{},
			"/ziraat": commands.ZiraatCommand{},
			"/deniz":  commands.DenizCommand{},
			"/iban":   commands.IbanCommand{},
		},
	}

	// Update HelpCommand with the commands list
	bot.commands["/help"] = commands.HelpCommand{Commands: bot.commands}

	return bot, nil
}

func (b *Bot) Start() {
	// b.api.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		cmd := strings.Split(update.Message.Text, " ")[0]
		args := strings.Fields(update.Message.Text)[1:]

		logMsg := fmt.Sprintf("[%s, id: %d] %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)
		log.Print(logMsg)

		if command, found := b.commands[cmd]; found {
			msg.Text = command.Execute(args)
		} else {
			msg.Text = "Неизвестная команда. Используйте /help для получения списка команд."
		}

		b.api.Send(msg)
	}
}
