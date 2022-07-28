package main

import (
	"fmt"
	"sort"
)

//给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
//假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
//你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
//示例 1：
//输入：nums = [1,3,4,2,2]
//输出：2
//示例 2：
//输入：nums = [3,1,3,4,2]
//输出：3
// 解法二 二分搜索
func findDuplicate(nums []int) int {
	n := len(nums)
	left, right := 1, n - 1
	ans := -1
	for left <= right {
		mid := (left + right) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}
	}
	return ans
}

// 解法三
func findDuplicate1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	diff := -1
	for i := 0; i < len(nums); i++ {
		if nums[i]-i-1 >= diff {
			diff = nums[i] - i - 1
		} else {
			return nums[i]
		}
	}
	return 0
}
// 解法一 快慢指针
func findDuplicate2(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fmt.Println("slow",slow)
		fast = nums[nums[fast]]
		fmt.Println("fast",fast)
	}
	walker := 0
	for walker != slow {
		walker = nums[walker]
		slow = nums[slow]
	}
	return walker
}

func main() {
	nums := []int{1,3,4,2,2}
	fmt.Println(findDuplicate2(nums))
}