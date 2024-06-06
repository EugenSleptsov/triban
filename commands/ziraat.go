package commands

import (
	"github.com/EugenSleptsov/triban/utils"
	"log"
	"strings"
)

type ZiraatCommand struct{}

func (z ZiraatCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Использование: /ziraat iban"
	}
	iban := strings.Join(args, " ")

	ibanData := utils.GetDataFromIban(iban)

	log.Print(ibanData)

	if ibanData.Error != "" {
		return ibanData.Error
	}

	if ibanData.BankCode != "10" {
		return "IBAN не относится к Ziraat Bankasi"
	}

	if ibanData.ClientNumber == "" {
		return "Номер клиента не найден"
	}

	return "Номер клиента: " + ibanData.ClientNumber
}

func (z ZiraatCommand) Description() string {
	return "Высчитывает номер клиента из IBAN Ziraat bankasi"
}
