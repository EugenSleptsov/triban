package commands

import "github.com/EugenSleptsov/triban/utils"

type DenizCommand struct{}

func (d DenizCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Usage: /deniz iban"
	}
	iban := args[0]
	return utils.CalculateAccountNumberFromIban(iban)
}

func (d DenizCommand) Description() string {
	return "Calculates account number from IBAN"
}
