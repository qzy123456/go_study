package main

import "fmt"

//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
// 解法一 ，投票法，时间复杂度 O(n) 空间复杂度 O(1),
//让每一个多数元素和其他任意元素抵消，那么最后剩下来的还是多数元素。
func majorityElement(nums []int) int {
	res   := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res   = nums[i]
			count = 1
		} else {
			if nums[i] == res {
				count++
			} else {
				count--
			}
		}
	}
	return res
}

// 解法二 时间复杂度 O(n) 空间复杂度 O(n)
func majorityElement1(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

func main()  {
	nums := []int{1,2,1,2,1}
	fmt.Println(majorityElement(nums))
	fmt.Println(majorityElement1(nums))
}
