package main

import "fmt"

func main() {
	/*
	 * Scrivere un programma che dato un numero in ingresso dall'utente
	 * verifichi se esso è pari o dispari
	 */
	
	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
