package main

import "fmt"

//有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
//现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 
//这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
//求所能获得硬币的最大数量。
//示例 1：
//输入：nums = [3,1,5,8]
//输出：167
//解释：
//nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
//coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
func maxCoins(nums []int) int {
	//因为会越界，提前给数组前后给一个数字
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	// 13158
	//初始化dp
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums))
	}

	for i := len(nums) - 3; i >= 0; i-- {
		for j := i + 2; j < len(nums); j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], nums[k]*nums[i]*nums[j]+dp[i][k]+dp[k][j])
			}
		}
	}
	return dp[0][len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}
