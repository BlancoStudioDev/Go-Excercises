package main

import (
	"fmt"
	"os"
)

func main() {
	var chars [5]byte
	_, err := fmt.Scanf("%c%c%c%c%c", &chars[0], &chars[1], &chars[2], &chars[3], &chars[4])
	if err != nil {
		fmt.Println("Errore durante la lettura:", err)
		return
	}

	max := chars[0]
	for _, c := range chars {
		if c > max {
			max = c
		}
	}

	fmt.Printf("%c", max)
}