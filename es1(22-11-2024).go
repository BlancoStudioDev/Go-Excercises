package main

import "fmt"

func incrementa(p *int){
	(*p)++
}

func main(){
	var x int
	x = 7
	incrementa(&x)
	fmt.Println(x)
}