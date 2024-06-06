package commands

import "fmt"

type HelpCommand struct {
	Commands map[string]Command
}

func (h HelpCommand) Execute(args []string) string {
	helpText := "Доступные команды:\n"
	for cmd, handler := range h.Commands {
		helpText += fmt.Sprintf("%s — %s\n", cmd, handler.Description())
	}
	return helpText
}

func (h HelpCommand) Description() string {
	return "Показывает список доступных команд"
}
