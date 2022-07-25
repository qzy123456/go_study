package main

import "fmt"

//动态规划
//动态规划方程：dp[n] = MAX( dp[n-1], dp[n-2] + num)
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
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

func rob2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length > 1 {
		// 如果第一家大于第二家，那么就偷第一家，否则就偷第二家
		if nums[1] < nums[0] {
			nums[1] = nums[0]
		}
		if length > 2 {
			for i := 2; i < len(nums); i++ {
				// 偷第i家
				if nums[i]+nums[i-2] >= nums[i-1] {
					nums[i] += nums[i-2]
				} else {
					// 不偷第i家
					nums[i] = nums[i-1]
				}
			}
		}
	}
	return nums[length-1]
}

func main() {
	arr  := []int{1, 2, 3, 1}
	arr2 := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(arr))
	fmt.Println(rob(arr2))
	fmt.Println(rob2(arr2))
}
