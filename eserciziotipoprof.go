//scrivere una funzione che date due slice di interi x e y restituisce una nuova slice che contiene
//solo gli elementi di x che non compaiono in y

package main

import "fmt"

func cerca(x []int, y []int) []int{
	var z []int
	for i:=0; i<len(x); i++{
		for j:=0; j<len(y); j++{
			if x[i] == y[j]{
				z = append(z, x[i])
				break
			}
		}
	}
	return z
}

func main () {
	var x []int
	var y []int
	x = []int{3,7,2,3,1}
	y = []int{3,5,3,3,3,2}
	fmt.Println(cerca(x,y))
}