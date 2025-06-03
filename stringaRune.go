package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Print("Inserisci una stringa: ")
	fmt.Scanln(&str)

	hasDigit := false
	for _, r := range str {
		if unicode.IsDigit(r) {
			hasDigit = true
			break
		}
	}

	if hasDigit {
		fmt.Println("ha cifre")
	} else {
		fmt.Println("nessuna cifra")
	}
}