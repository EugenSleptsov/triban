package utils

import "strings"

func FormatIban(iban string) string {
	// Format IBAN to human readable format
	// XXXX XXXX XXXX XXXX XXXX XXXX XXXX XXXX

	// Remove spaces
	iban = strings.ReplaceAll(iban, " ", "")

	// Upper case
	iban = strings.ToUpper(iban)

	// Create a slice to store the formatted IBAN parts
	var formatted []string

	// Iterate over the IBAN and collect chunks of four characters
	for i := 0; i < len(iban); i += 4 {
		end := i + 4
		if end > len(iban) {
			end = len(iban)
		}
		formatted = append(formatted, iban[i:end])
	}

	// Join the chunks with spaces
	return strings.Join(formatted, " ")
}
