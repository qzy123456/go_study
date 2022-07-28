package main

import (
	"fmt"
	"sort"
)

// 输入 nums =[10,9,2,5,3,7,101,18]
//最长递增子序列是 2,3,7,101 所以长度是4

//动态规划+sort.SearchInts()
func lengthOfLIS(nums []int) int {
	var dp []int
	for _, num := range nums {
		i := sort.SearchInts(dp, num) //min_index
		fmt.Println(i)
		if i == len(dp) {
			dp = append(dp, num) //num 比dp中的元素都大 应该加到末尾
		} else {
			dp[i] = num //num加到dp【i】处
		}
	}
	//fmt.Println(dp)
	return len(dp)
}

//动态规划 + 二分法
func lengthOfLIS2(nums []int) int {
	var dp []int
	for _, v := range nums {
		//如果v比dp的最后一个元素还大，说明v放到dp后还是符合严格递增的
		if len(dp) == 0 || v > dp[len(dp)-1] {
			dp = append(dp, v)
		} else {
			var left = 0
			var right = len(dp)
			for left < right {
				mid := left + (right-left)/2
				if dp[mid] < v {
					left = mid + 1
				} else {
					right = mid
				}
			}
			//此时left == right 哪个都一样
			dp[left] = v
		}
	}
	//fmt.Println(dp)
	return len(dp)
}
//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
//示例 1：
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//示例 2：
//输入：nums = [0,1,0,3,2,3]
//输出：4
//示例 3：
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
func lengthOfLIS3(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}
	return res
}
func max(a ,b int)int  {
	if a > b {
		return a
	}
	return b
}
func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//fmt.Println(lengthOfLIS(nums))
	fmt.Println(lengthOfLIS2(nums))
	fmt.Println(lengthOfLIS3(nums))
}
