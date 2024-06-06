package commands

import (
	"fmt"
	"github.com/EugenSleptsov/triban/utils"
)

type IninalCommand struct{}

func (i IninalCommand) Execute(args []string) string {
	if len(args) < 1 {
		return "Использование: /ininal accountNumber"
	}
	accountNumber := args[0]

	if len(accountNumber) != 13 {
		return "Номер аккаунта должен состоять из 13 цифр"
	}

	// check only digits
	for _, r := range accountNumber {
		if r < '0' || r > '9' {
			return "Номер аккаунта должен состоять только из цифр"
		}
	}

	return calculateIninalIban(accountNumber)
}

func (i IninalCommand) Description() string {
	return "Рассчитывает IBAN сервиса Ininal по номеру аккаунта (для тех у кого нет ininal plus)"
}

func calculateIninalIban(accountNumber string) string {
	//98 - (83200001625293556222292700 mod 97)

	// формируем число по маске 8320000аккаунт292700
	tmpNumber := "8320000" + accountNumber + "292700"

	// вычисляем остаток от деления на 97
	remainder := utils.Mod97(tmpNumber)

	// вычитаем остаток из 98
	checksum := 98 - remainder

	iban := fmt.Sprintf("TR%02d008320000%s", checksum, accountNumber)

	return utils.FormatIban(iban)
}
