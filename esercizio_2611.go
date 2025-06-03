package main

import "fmt"

func main(){
	str := "radar"
	fmt.Println(isPal(str))
}

func isPal(str string) bool{
	if len(str) <= 1{
		return true
	}
	if str[0] != str[len(str) - 1]{
		return false
	}
	return isPal(str[1:len(str)-1])
}

/*
 * per convertire una stringa in una runa:
 * stringa = []rune(stringa)
 */