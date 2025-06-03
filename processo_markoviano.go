package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"time"
)

const lungChiave int = 4
const lungValore int = 4

func leggi() map[string][]string {
	risultato := make(map[string][]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		for i := 0; i < len(riga)-lungChiave-lungValore; i++ {
			chiave := riga[i : i+lungChiave]
			valore := riga[i+lungChiave : i+lungChiave+lungValore]
			risultato[chiave] = append(risultato[chiave], valore)
		}
	}
	return risultato
}

func chiaveACaso(modello map[string][]string) string {
	// Seleziona una chiave casuale
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(modello))
	for k := range modello {
		if i == 0 {
			return k
		}
		i--
	}
	return ""
}

func genera(n int, modello map[string][]string, partenza string) string {
	risultato := partenza
	for len(risultato) < n {
		chiave := risultato[len(risultato)-lungChiave:]
		valore, ok := modello[chiave]
		if ok {
			// Se la chiave è presente, scegli un valore casuale
			risultato += valore[rand.Intn(len(valore))]
		} else {
			// Se la chiave non è presente, scegli una chiave casuale
			risultato += chiaveACaso(modello)
		}
	}
	return risultato
}

func main() {
	// Leggi il modello
	modello := leggi()
	// Converte il primo argomento della riga di comando in un intero
	n, _ := strconv.Atoi(os.Args[1])
	// Stampa la stringa generata
	fmt.Println(genera(n, modello, chiaveACaso(modello)))
}
