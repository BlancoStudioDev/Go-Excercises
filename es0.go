package main

import "fmt"

func main(){
	var x string
	fmt.Println("immetti valore: ")
	fmt.Scan(&x)
	fmt.Println(x)
	fmt.Println(string(x[0]+1))
	fmt.Println(string(x[0]-1))
	if x[0]>65 || x[0]<76{
		fmt.Println("A-L")
	} else {
		fmt.Println("altro")
	}
	var stringa string
	fmt.Println("immetti una parola")
	fmt.Scan(&stringa)
	for i:=0; i<len(stringa); i++{
		fmt.Println(string(stringa[i])) 
	}
}