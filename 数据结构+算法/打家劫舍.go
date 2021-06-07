package main

import "fmt"

//动态规划
//动态规划方程：dp[n] = MAX( dp[n-1], dp[n-2] + num)
func rob(nums []int) int {
	preMAX := 0
	currMAX := 0
	for i := 0; i < len(nums); i++ {
		temp := currMAX
		if currMAX <= preMAX + nums[i] {
			currMAX = preMAX + nums[i]
		}
		preMAX = temp
	}
	return currMAX
}

func main() {
	arr  := []int{1, 2, 3, 1}
	arr2 := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(arr))
	fmt.Println(rob(arr2))
}
