package main

import (
	"fmt"
	"strings"
)

func main() {
	var char rune
	var str string
	fmt.Print("Inserisci un carattere: ")
	fmt.Scan(&char)
	fmt.Print("Inserisci una stringa: ")
	fmt.Scan(&str)

	index := strings.IndexRune(str, char)
	fmt.Println(index)
}