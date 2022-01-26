package main

import "fmt"

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。
//返回滑动窗口中的最大值所构成的数组。
func maxSlidingWindow(nums []int, k int) []int {
	l1 := len(nums) //数组长度
	index := 0      //起始下标
	ret := make([]int, 0) //返回值
	for index < l1 {
		m := nums[index]
		if index > l1 - k {
			break
		}
		for j := index + 1; j < index + k; j++ {
			if m < nums[j] {
				m = nums[j]
			}
		}
		ret = append(ret,m)
		index++
	}
	return ret
}
func main() {
	//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3 输出: [3,3,5,5,6,7] 解释:
   //[1 3 -1] -3 5 3 6 7  输出 3
	//
	//1 [3 -1 -3] 5 3 6 7  输出 3
	//
	//1 3 [-1 -3 5] 3 6 7  输出 5
	//
	//1 3 -1 [-3 5 3] 6 7  输出 5
	//
	//1 3 -1 -3 [5 3 6] 7  输出 6
	//
	//1 3 -1 -3 5 [3 6 7]  输出 7
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	fmt.Println(maxSlidingWindow(nums,k))
}
