package commands

import "strings"

type ZiraatCommand struct{}

func (z ZiraatCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Usage: /ziraat iban"
	}
	iban := strings.Join(args, " ")
	return iban
}

func (z ZiraatCommand) Description() string {
	return "Высчитывает номер клиента из IBAN Ziraat bankasi"
}
