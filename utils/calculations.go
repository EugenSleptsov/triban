package utils

import "fmt"

func CalculateIninalIban(accountNumber string) string {
	// Dummy implementation
	return fmt.Sprintf("IBAN for account number %s is: TRXX XXXX XXXX XXXX XXXX XXXX XX", accountNumber)
}

func CalculateAccountNumberFromIban(iban string) string {
	// Dummy implementation
	return fmt.Sprintf("Account number for IBAN %s is: XXXXXXXXXX", iban)
}
