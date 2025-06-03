package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Inserisci una stringa: ")
	fmt.Scanln(&str)

	if len(str) > 0 {
		prev := str[0]
		fmt.Print(string(prev))
	}

	for i := 1; i < len(str); i++ {
		curr := str[i]
		if curr > prev {
			fmt.Print("+")
		} else if curr < prev {
			fmt.Print("-")
		} else {
			fmt.Print("=")
		}
		fmt.Print(string(curr))
		prev = curr
	}

	fmt.Println()
}