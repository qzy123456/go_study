package main

import "fmt"

//其实就是位1的个数差异 ，1011101 与 1001001 之间的汉明距离是 2。
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