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
			"/usd":    &commands.UsdCommand{},
		},
	}

	// Update HelpCommand with the commands list
	bot.commands["/help"] = commands.HelpCommand{Commands: bot.commands}

	return bot, nil
}

func (b *Bot) Start() {
	log.Println("Bot started")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		cmd := strings.Split(update.Message.Text, " ")[0]
		args := strings.Fields(update.Message.Text)[1:]

		logMsg := fmt.Sprintf("[%s, id: %d] %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)
		log.Print(logMsg)

		if command, found := b.commands[cmd]; found {
			b.SendMarkdown(update.Message.Chat.ID, command.Execute(args))
		} else {
			b.Send(update.Message.Chat.ID, "Неизвестная команда. Используйте /help для получения списка команд.", false)
		}
	}
}

func (b *Bot) SendMarkdown(chatID int64, text string) {
	b.Send(chatID, text, true)
}

func (b *Bot) Send(chatID int64, text string, isMarkdown bool) {
	msg := tgbotapi.NewMessage(chatID, text)
	if isMarkdown {
		msg.ParseMode = "MarkdownV2"
		msg.Text = escapeMarkdownV2(msg.Text)
	}
	_, err := b.api.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func escapeMarkdownV2(text string) string {
	charsToEscape := []string{"_", "*", "[", "]", "(", ")", "~", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	for _, char := range charsToEscape {
		text = strings.ReplaceAll(text, char, "\\"+char)
	}
	return text
}
