package commands

import "github.com/EugenSleptsov/triban/utils"

type ZiraatCommand struct{}

func (z ZiraatCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Usage: /ziraat iban"
	}
	iban := args[0]
	return utils.CalculateAccountNumberFromIban(iban)
}

func (z ZiraatCommand) Description() string {
	return "Calculates account number from IBAN"
}
