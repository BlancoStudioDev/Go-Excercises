package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func attendi(n_seconds float64) {
	time.Sleep(time.Duration(n_seconds) * time.Second)
}

func cancella() {
	cmd := exec.Command("cmd", "/C", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func stampaMatrice(flag int) {
	cols, rows := 8, 8
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	valore := 7
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if flag == 0{
				if i == j {
					matrix[i][j] = 2
				} else if j == valore {
					matrix[i][j] = 1
				} else if (i <= j) && (j <= valore) && (i <= rows/2) {
					matrix[i][j] = 3
				} else {
					matrix[i][j] = 0
				}
			} else if flag == 1{
				if i == j {
					matrix[i][j] = 2
				} else if j == valore {
					matrix[i][j] = 1
				} else if (i <= j) && (j <= valore) && (i <= rows/2 && i>=1) ||  i == 7 {
					matrix[i][j] = 3
				} else {
					matrix[i][j] = 0
				}
			} else if flag == 2{
				if i == j {
					matrix[i][j] = 2
				} else if j == valore {
					matrix[i][j] = 1
				} else if (i >= j) && (j >= valore) &&  (i == 7 || i == 6) || (i <= j) && (j <= valore) &&  (i == 3 || i == 2)  {
					matrix[i][j] = 3
				} else {
					matrix[i][j] = 0
				}
			} else if flag == 3{
				if i == j {
					matrix[i][j] = 2
				} else if j == valore {
					matrix[i][j] = 1
				} else if (i > j) && (j > valore) && (i == 7 || i == 6 || i == 5)  {
					matrix[i][j] = 3
				} else {
					matrix[i][j] = 0
				}
			}
		}
		valore--
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				fmt.Print(" ")
			} else if matrix[i][j] == 2 {
				fmt.Print("\\")
			} else if matrix[i][j] == 1 {
				fmt.Print("/")
			} else if matrix[i][j] == 3 {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func main() {
	for i:=0; i<4; i++{
		stampaMatrice(i)
		attendi(1)
		cancella()
	}
}
