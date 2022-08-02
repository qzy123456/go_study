package main

import "fmt"

func main() {
	yanghuisanjiao(12)
}

func yanghuisanjiao(rows int) {
	for i := 0; i < rows; i++ {
		number := 1
		for k := 0; k < rows-i; k++ {
			fmt.Print("  ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%4d", number)
			number = number * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
