/*
 * dichiarazione: var x []int
 * creazione della slice -> x = make([]tipo di slice, lunghezza sotto forma di espressione(come variabili))
 * le slice posso allungarle di quanto voglio, per farlo si usa una funzione di libreria che si chiama append:
 * x = append(x, element, element1, ...)	
 */

package main

import "fmt"

func main(){
	var x []int
	// x = make([]int, 10) o questo 
	for {
		var h int
		fmt.Scan(&h)
		if h == 0 {
			break
		} else {
			x = append(x, h)
		}
	}
	fmt.Println(x)
}