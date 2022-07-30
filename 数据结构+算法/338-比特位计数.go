package main

import "fmt"

//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，
//返回一个长度为 n + 1 的数组 ans 作为答案。
//示例 1：
//输入：n = 2
//输出：[0,1,1]
//解释：
//0 --> 0
//1 --> 1
//2 --> 10
//示例 2：
//输入：n = 5
//输出：[0,1,1,2,1,2]
//解释：
//0 --> 0
//1 --> 1
//2 --> 10
//3 --> 11
//4 --> 100
//5 --> 101
//  X&1==1or==0，可以用 X&1 判断奇偶性，X&1>0 即奇数。
//  X = X & (X-1) 清零最低位的1
//  X & -X => 得到最低位的1
//  X&~X=>0
func countBits(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] += bits[i&(i-1)] + 1
	}
	return bits
}

func main() {
	fmt.Println(countBits(5))
}

