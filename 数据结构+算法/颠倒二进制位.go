package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var res uint32  = 0
	for  i := 0; i < 32; i++ {
		//res先往左移一位，把最后一个位置空出来，
		//用来存放n的最后一位数字
		res <<= 1
		//res加上n的最后一位数字
		res += num & 1
		//n往右移一位，把最后一位数字去掉
		num >>= 1
	}
	return res
}

func main() {
	fmt.Println(reverseBits(12))
}


