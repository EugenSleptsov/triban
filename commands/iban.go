package commands

import (
	"github.com/EugenSleptsov/triban/utils"
	"strings"
)

type IbanCommand struct{}

func (i IbanCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Использование: /iban IBAN"
	}
	iban := strings.Join(args, "")
	ibanData := utils.GetDataFromIban(iban)

	str := "IBAN: " + utils.FormatIban(ibanData.Iban) + "\n"
	if ibanData.Error != "" {
		return str + ibanData.Error
	}

	str += "Страна: " + ibanData.CountryCode + "\n"
	str += "Контрольная сумма: " + ibanData.Checksum + "\n"
	str += "BBAN: " + ibanData.Bban + "\n"

	if ibanData.Bank != "" {
		str += "Банк: " + ibanData.Bank + "\n"
	} else if ibanData.Fintech != "" {
		str += "Финансовая организация: " + ibanData.Fintech + "\n"
	}

	if ibanData.AccountNumber != "" {
		str += "Номер счета: " + ibanData.AccountNumber + "\n"
	}
	if ibanData.ClientNumber != "" {
		str += "Номер клиента: " + ibanData.ClientNumber + "\n"
	}
	if ibanData.Description != "" {
		str += ibanData.Description + "\n"
	}

	return str
}

func (i IbanCommand) Description() string {
	return "Предоставляет информацию, которую можно получить из IBAN"
}
