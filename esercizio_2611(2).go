package main

import "fmt"

func main(){
	x := 1110001
	fmt.Println(quantiuni(x))
}

func quantiuni(x int)int{
	if x<=9{
		if x == 1{
			return 1
		} else {
			return 0
		}
	}
	if x%10 == 1{
		return 1 + quantiuni(x%10)
	} else {
		return quantiuni(x%10)
	}
}