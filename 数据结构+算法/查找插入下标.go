package main

import "fmt"

//查找要插入的下标
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}
//暴力解决问题，顺序查找
func searchInsert1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		} else {
			if i == len(nums) - 1 {
				return i + 1
			}
		}
	}
	return len(nums) - 1
}
//二分法
//可以想到这是一道类似于在0 0 0 0 1 1 1 1数列中找到第一个1的问题
//那么0和1的意义要明确清楚
//0代表target大于当前值
//1代表target小于等于当前值
//但是还有一种特殊情况我们需要判断
//当target大于所有的当前值时
//left会越界，所以在循环条件中要加上left < len(nums)
func searchInsert2(nums []int, target int) int {
	left := 0
	right := len(nums)
	mid := 0
	for left < right && left < len(nums) {  //防止越界
		mid = (left + right) >> 1   //取一个中位
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}


func main() {
	nums :=[]int{1,2,3,4,5}
	fmt.Println(searchInsert(nums,3))
	fmt.Println(searchInsert1(nums,3))
	fmt.Println(searchInsert2(nums,3))
}