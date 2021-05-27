package main

import (
	"fmt"
	"math"
)

//从一个未排序的数组里找到连续递增的长度

//贪心算法
func findLength(nums []int)int  {
	var start = 0
	var maxLength = 0

	for i:=1;i<len(nums) ;i++  {
		if nums[i] <= nums[i-1]{
           	start = i
		}
		maxLength = int(math.Max(float64(maxLength),float64(i-start+1)))
	}
	return maxLength
}
//双指针
func findLength2(nums []int)int  {
	var low = 0
	var height = 1
	var leng  = len(nums)
	var maxLength  = 0
	for low<leng && height<leng  {
		if nums[height] <= nums[height-1]{
			low = height
		}
		maxLength = int(math.Max(float64(maxLength),float64(height-low+1)))
		height++
	}
	return maxLength
}
func main() {
	nums := []int{1,2,3,2,4,5}
	fmt.Println(findLength(nums))
	fmt.Println(findLength2(nums))
}