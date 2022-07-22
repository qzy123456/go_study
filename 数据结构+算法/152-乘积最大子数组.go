package main

import "fmt"


//给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
func maxProduct(nums []int) int {
	//最小值，最大值 默认取第一个数
	minimum, maximum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0{
			maximum, minimum = minimum, maximum
		}
		minimum = min(nums[i], minimum * nums[i])
		maximum = max(nums[i], maximum * nums[i])
		res = max(res, maximum)
	}
	return res
}

func max(a ,b int)int{
	if a > b {
		return a
	}
	return b
}
func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}

func main() {
	nums := []int{2,3,-2,4}
	fmt.Println(maxProduct(nums))
}
