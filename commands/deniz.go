package commands

import (
	"github.com/EugenSleptsov/triban/utils"
	"strings"
)

type DenizCommand struct{}

func (d DenizCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Использование: /deniz iban"
	}
	iban := strings.Join(args, " ")

	ibanData := utils.GetDataFromIban(iban)

	if ibanData.Error != "" {
		return ibanData.Error
	}

	if ibanData.BankCode != "134" {
		return "IBAN не относится к Denizbank"
	}

	if ibanData.ClientNumber == "" {
		return "Номер клиента не найден"
	}

	return "Номер клиента: " + ibanData.ClientNumber
}

func (d DenizCommand) Description() string {
	return "Высчитывает номер клиента из IBAN Denizbank"
}
