package main

import (
	"fmt"
	"sort"
)

//581. 最短无序连续子数组
//给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//请你找出符合题意的 最短 子数组，并输出它的长度。
//示例 1：
//
//输入：nums = [2,6,4,8,10,9,15]
//输出：5
//解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//示例 2：
//
//输入：nums = [1,2,3,4]
//输出：0
//示例 3：
//
//输入：nums = [1]
//输出：0
func findUnsortedSubarray(nums []int) int {

	num2 := make([]int, len(nums))
	copy(num2, nums)
	//复制一个数组，并排序
	sort.Ints(num2)
	//左，右指针，分别记录 最大 最小
	lo, hi := -1, -1

	for key, value := range nums {
		//如果是第一位，说明是最小的。因为nums2排过序了
		if lo == -1 && value != num2[key] {
			lo = key
		}
		//最大的连续改变
		if value != num2[key] {
			hi = key
		}
	}
	//如果已经是排过序的数组
	if hi == lo {
		return 0
	}
	return hi - lo + 1
}

func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(nums))
}
