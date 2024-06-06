package commands

import "strings"

type DenizCommand struct{}

func (d DenizCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Использование: /deniz iban"
	}
	iban := strings.Join(args, " ")
	return iban
}

func (d DenizCommand) Description() string {
	return "Высчитывает номер клиента из IBAN Denizbank"
}
