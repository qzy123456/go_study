package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func numTrees2(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res = res * 2 * (2 * i + 1) / (i + 2)
	}
	return res
}

func main() {
	fmt.Println(numTrees(2))
	fmt.Println(numTrees2(2))
}
