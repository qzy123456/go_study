package main

import "fmt"

func hammingWeight(num uint32) int {
	temp := 0
	for ;num != 0;temp++{
		num = num & (num - 1)
	}
	return temp
}

func main()  {
	var n uint32 = 3
	fmt.Println(hammingWeight(n))
}