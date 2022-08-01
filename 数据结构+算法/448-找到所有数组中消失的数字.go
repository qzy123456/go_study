package main

import "fmt"

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，
//并以数组的形式返回结果。
//示例 1：
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[5,6]
//示例 2：
//
//输入：nums = [1,1]
//输出：[2]

func findDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	fmt.Println(nums) //[12 19 18 15 8 2 11 9]
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}

func main() {
	nums := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDisappearedNumbers(nums))
}

