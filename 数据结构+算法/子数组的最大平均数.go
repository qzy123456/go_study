package main

import (
	"fmt"
	"math"
)

//给定一个整数数组，找出平均数最大且长度为K的下标连续的子数组，
//并输出该平均数。
//如[1,12,-5,-6,50,3], k=4
//返回 12.75
//滑动窗口方法
func findMaxAverage(nums []int, k int)float64  {
	var sum  = 0
	var n  = len(nums)
	//使用第一个窗口的值来初始化sum
	for i:=0;i<k ;i++  {
		sum += nums[i]
	}
	//前4位sum为2
	var max  = sum
	//开始向右滑动，i为右边界
	for i:=k;i<n ;i++  {
		// i-k 为左边界
		sum = sum  + nums[i] - nums[i-k]
		fmt.Println(nums[i-k],nums[i],sum,max)
		//1 50 51 2
		//12 3 42 51
		max = int(math.Max(float64(sum),float64(max)))
	}
	return float64(max) /float64(k)
}
//双指针版滑动窗口
func findMaxAverage2(nums []int,k int)float64  {
	var low  = 0  //左边界减去的旧窗口
	var high  = k //右边界的新增窗口
	var n  = len(nums)
	var sum  = 0
	//使用第一个窗口的值来初始化sum
	for i:=0;i<k ;i++  {
		sum += nums[i]
	}
	var max  = sum
	for high < n{
		sum = sum + nums[high] - nums[low]
		max = int(math.Max(float64(sum),float64(max)))
		high++
		low++
	}
	return float64(max) / float64(k)
}

func main() {
	nums := []int{1,12,-5,-6,50,3}
	fmt.Println(findMaxAverage(nums,4))
	fmt.Println(findMaxAverage2(nums,4))
}