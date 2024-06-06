package utils

import (
	"github.com/EugenSleptsov/triban/consts"
	"strings"
)

type IbanData struct {
	Iban          string
	CountryCode   string
	Checksum      string
	Bban          string
	BankCode      string
	Bank          string
	Fintech       string
	AccountNumber string
	ClientNumber  string
	Error         string
	Description   string
}

func GetDataFromIban(iban string) IbanData {
	ibanData := IbanData{Iban: iban}

	iban = strings.ReplaceAll(iban, " ", "")
	if !CheckIban(iban) {
		ibanData.Error = "IBAN выглядит неверным (ошибка в контрольной сумме), проверьте правильность и введите правильный IBAN"
		return ibanData
	}

	if len(iban) < 5 {
		ibanData.Error = "IBAN слишком короткий"
		return ibanData
	}

	ibanData.CountryCode = iban[0:2]
	ibanData.Checksum = iban[2:4]
	ibanData.Bban = iban[4:]

	if ibanData.CountryCode == "TR" {
		if len(iban) != 26 {
			ibanData.Error = "IBAN относится к Турции, но длина IBAN не равна 26 символам"
			return ibanData
		}

		// Trim слева по нулям
		ibanData.BankCode = strings.TrimLeft(ibanData.Bban[0:5], "0")

		if consts.BankCodes[ibanData.BankCode] != "" {
			ibanData.Bank = consts.BankCodes[ibanData.BankCode]
			if ibanData.BankCode == "15" {
				ibanData.AccountNumber = ibanData.Bban
			} else if ibanData.BankCode == "10" {
				ibanData.ClientNumber = ibanData.Bban[8:17]
			} else if ibanData.BankCode == "134" {
				ibanData.ClientNumber = ibanData.Bban[9:17]
			}
		} else if consts.PaymentCodes[ibanData.BankCode] != "" {
			ibanData.Fintech = consts.PaymentCodes[ibanData.BankCode]

			if ibanData.BankCode == "832" {
				ibanData.AccountNumber = ibanData.Bban[9:]
			} else if ibanData.BankCode == "829" {
				ibanData.AccountNumber = ibanData.Bban[12:]
				ibanData.ClientNumber = ibanData.Bban[12:]
			}
		} else {
			ibanData.Bank = "Неизвестный банк или финансовая организация"
		}
	}

	return ibanData
}
