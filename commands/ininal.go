package commands

import "github.com/EugenSleptsov/triban/utils"

type IninalCommand struct{}

func (i IninalCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Usage: /ininal account_number"
	}
	accountNumber := args[0]
	return utils.CalculateIninalIban(accountNumber)
}

func (i IninalCommand) Description() string {
	return "Calculates IBAN from account number"
}
