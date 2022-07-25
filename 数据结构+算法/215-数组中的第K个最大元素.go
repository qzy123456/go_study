package main

import (
	"fmt"
	"sort"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
// 解法一 排序，排序的方法反而速度是最快的
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//解法2，快排
func findKthLargest(nums []int, k int) int {
	// 传入元素进行查询
	top(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func top(nums []int,k,start,stop int) {
	// 若数组长度为 1，则直接返回
	if start >= stop {
		return
	}
	// 进行单次快排，得出第一个对比数的索引
	index := parition(nums,start,stop)
	// 索引正好为 k,则 k 为第 k 个最大的元素
	if index == k {
		return
	}

	// 如果索引比k小，那么将索引右部进行单次快排
	if index < k {
		top(nums,k,index+1,stop)
		// 如果索引比k大，那么将索引左部进行单次快排
	} else if index > k {
		top(nums,k,start,index-1)
	}

}

// parition 单次快排
func parition(nums []int, start, stop int) int {
	if start >= stop{
		return -1
	}

	//将第一个数作为快排基准数
	pivot := nums[start]
	// 获得两个函数内部变量，作为快排左右指针
	l, r := start, stop
	// 如果左指针小于右指针，则继续排序
	for l < r{
		// 如果左指针小于右指针且右指针元素大于基准数
		// 则该元素满足快排原则，右指针左移后再次判断
		for l < r && nums[r] >= pivot{
			r--
		}
		// 如果左指针小于右指针且左指针元素小于基准数
		// 则该元素满足快排原则，左指针右移后再次判断
		for l < r && nums[l] <= pivot{
			l++
		}
		// 由于此时左右指针停止，则左指针元素大于基准数、右指针元素小于基准数
		// 此时交换左右指针元素
		nums[r],nums[l] = nums[l],nums[r]
	}

	// 循环退出、证明左右指针相遇，此时指针停留位置小于等于基准数
	// 将基准数与指针数交换、此时基准数的位置即为指针位置
	nums[l], nums[start] = nums[start], nums[l]
	return l
}

func main() {
	nums := []int{1,2,3,4,5,6}
	k := 2
	fmt.Println(findKthLargest(nums,k))
}