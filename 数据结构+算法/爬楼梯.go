package main

import "fmt"

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	first := 1
	second := 2
	sum := 0
	for n > 2{
		sum = first + second
		first = second
		second = sum
		n--
	}
	return sum
}
func climbStairs1(n int) int {
	if n <= 1 {
		return 1
	}
	if n < 3 {
		return n
	}
	return climbStairs1(n - 1) + climbStairs1(n - 2)
}
func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
//滚动数组
func climbStairs3(n int) int {
	dp := [2]int{1, 1}

	for i := 2; i <= n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}

	return dp[n%2]
}
func main() {
   n := climbStairs(5)
   fmt.Println(n)
   fmt.Println(climbStairs1(5))
   fmt.Println(climbStairs2(5))
   fmt.Println(climbStairs3(5))

}
