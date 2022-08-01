package main

import "fmt"

//416. 分割等和子集
//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//示例 1：
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//示例 2：
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//给定一个非空的数组，其中所有的数字都是正整数。问是否可以将这个数组的元素分为两部分，使得每部分的数字和相等。
//这一题是典型的完全背包的题型。在 n 个物品中选出一定物品，完全填满 sum/2 的背包。
//F(n,mid) 代表将 n 个物品填满容量为 mid 的背包，
//状态转移方程为 F(i,mid) = F(i - 1,mid) || F(i - 1, mid - w[i])。当 i - 1 个物品就可以填满 mid
//这种情况满足题意。同时如果 i - 1 个物品不能填满背包，加上第 i 个物品以后恰好可以填满这个背包，也可以满足题意。
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	n   := len(nums)
	mid := sum/2
	fmt.Println(mid)
	dp  := make([]bool, sum/2+1)
	for i := 0; i <= mid; i++ {
		dp[i] = nums[0] == i
	}
	fmt.Println(dp)
	for i := 1; i < n; i++ {
		for j := mid; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[mid]
}
func main() {
	nums := []int{1,5,11,5}
	fmt.Println(canPartition(nums))
}