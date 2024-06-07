package utils

import (
	"fmt"
	"math/big"
	"strings"
	"unicode"
)

// CheckIban validates an IBAN
func CheckIban(iban string) bool {
	// Remove spaces and to upper case
	iban = strings.ToUpper(strings.ReplaceAll(iban, " ", ""))

	// Move the first four characters to the end
	iban = iban[4:] + iban[:4]

	// Replace each letter in the string with two digits
	var numericIban string
	for _, r := range iban {
		if unicode.IsLetter(r) {
			numericIban += fmt.Sprintf("%d", r-'A'+10)
		} else {
			numericIban += string(r)
		}
	}

	// Convert the numeric string to a big integer
	ibanInt := new(big.Int)
	ibanInt.SetString(numericIban, 10)

	// Perform the mod-97 operation
	return ibanInt.Mod(ibanInt, big.NewInt(97)).Int64() == 1
}

// Mod97 Остаток от деления на 97
func Mod97(number string) int {
	remainder := 0
	for _, r := range number {
		remainder = (remainder*10 + int(r-'0')) % 97
	}
	return remainder
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string with 4 digits
	return fmt.Sprintf("%.4f", input_num)
}
