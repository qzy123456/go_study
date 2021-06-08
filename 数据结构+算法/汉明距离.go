package main

import "fmt"

func hammingDistance(x int, y int) int {
	x = x^y
	var count = 0
	for x!=0{
		x = x & (x -1)
		count++
	}
	return count
}

func main()  {
	fmt.Println(hammingDistance(1,2))
}