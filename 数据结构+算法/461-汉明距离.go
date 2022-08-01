package main

import (
	"fmt"
	"math"
)

//其实就是位1的个数差异 ，1011101 与 1001001 之间的汉明距离是 2。
//输入：x = 1, y = 4
//输出：2
//解释：
//1   (0 0 0 1)
//4   (0 1 0 0)
//       ↑   ↑
//上面的箭头指出了对应二进制位不同的位置。
func hammingDistance(x int, y int) int {
	x = x ^ y
	var count = 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
//只支持正整数
func hammingDistance2(x int, y int) int {
	x = x ^ y
	return bitcount(x)
}
func bitcount(i int) int {
	count := 0
	i = int(math.Abs(float64(i)))
	for i > 0 {
		if i%2 == 1 {
			count++
		}
		i = i >> 1
	}

	return count
}
func main() {
	fmt.Println(hammingDistance(1, 2))
	fmt.Println(hammingDistance2(-1, 2))
}
