package main

import (
	"math"
	"fmt"
)

func main(){


	var num1, num2 int
	
	fmt.Print("Immetti il primo valore")
	fmt.Scan("%d", &num1)

	fmt.Print("Immetti il secondo valore")
        fmt.Scan("%d", &num2)

	fmt.Print(num1, num2)
}
