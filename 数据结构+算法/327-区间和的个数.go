package main

import "fmt"

func countRangeSum(nums []int, lower, upper int) int {
	n := len(nums)
	count := 0
	dp := make([][]int, n)
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int,n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = nums[i] //初始化
			} else {
				dp[i][j] = dp[i][j-1] + nums[j]
			}
			if dp[i][j] >= lower && dp[i][j] <= upper {
				count++
			}
		}
	}
	return count
}
func main() {
	nums := []int{-2,5,-1}
	fmt.Println(countRangeSum(nums,-2,2))
}