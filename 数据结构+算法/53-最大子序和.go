package main

import "fmt"

//给定一个整数数组nums，找到一个拥有最大连续和的连续子数组
//nums = [-2，1，-3，4，-1，2，1，-5，4] 连续子数组  4，-1，2，1的和最大为6
func maxSubArray(nums []int) int {
	var res = nums[0] //默认第一个数为最大和
	var sum = 0       //sum为前面nums【i】前面元素的最大和
	for i := 0; i < len(nums); i++ {
		//fmt.Println(sum + nums[i], nums[i])
		//-2 -2   -2 +0，-2 ==》-2
		//-1 1    -2+1 ，1  ==》1
		//-2 -3   -3+1，-3  ==》-2
		//2 4     -2+4，4   ==》4
		//3 -1    -1+4，-1  ==》3
		//5 2      3+2，2   ==》5
		//6 1
		//1 -5
		//5 4
		//6
		sum = max(sum + nums[i], nums[i])
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//动态规划解法
func maxSubArrayKMP(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
	fmt.Println(maxSubArrayKMP(nums))
}
