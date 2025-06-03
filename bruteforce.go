package main

import "fmt"

func main(){
	x := 192900
	for i:=0; i<9999999999; i++{
		j:=i-1
		if x == j{
			break
		}
		fmt.Println(i)
	}
}