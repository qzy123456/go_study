package main

import (
	"fmt"
)

//给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
//例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//示例 1：
//输入：n = 12
//输出：3
//解释：12 = 4 + 4 + 4
//示例 2：
//输入：n = 13
//输出：2
//解释：13 = 4 + 9
// 1  2    3      4    	5
// 1 1+1 1+1+1   2*2  2*2+1
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares(12))
}
