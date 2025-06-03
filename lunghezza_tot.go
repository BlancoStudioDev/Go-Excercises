package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var totLen int
	fmt.Print("Inserisci il valore di totLen: ")
	fmt.Scanln(&totLen)

	var totalLength int
	var concatStr strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		strLen := len(str)
		totalLength += strLen
		if totalLength >= totLen {
			break
		}
		concatStr.WriteString(str)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Errore durante la lettura:", err)
		return
	}

	fmt.Println("Lunghezza totale:", totalLength)
	fmt.Println("Concatenazione:", concatStr.String())
}