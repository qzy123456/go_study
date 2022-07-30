package main

import "fmt"

func rob3(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	//换一种思路，假如小偷走两条路线，这样就包含了首尾。最后返回最大的
	//2, 7, 9, 3, 1
	max1 := isRob(nums, 0, length-2)//2,7,9,3
	max2 := isRob(nums, 1, length-1)//7,9,3,1

	return max(max1, max2)
}
//下面的逻辑其实就是打家劫舍1
func isRob(nums []int, start, end int) int {
	//Rob := 0 //偷
	//nRob := 0 //不偷
	//for i := start; i <= end; i++ {
	//	tempMax := max(Rob, nRob)
	//	Rob = nums[i] + nRob
	//	nRob = tempMax
	//}
	//return max(Rob, nRob)
	//
	preMAX := 0
	currMAX := 0
	for i := start; i < end; i++ {
		temp := currMAX
		if currMAX <= preMAX + nums[i] {
			currMAX = preMAX + nums[i]
		}
		preMAX = temp
	}
	return currMAX
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	arr := []int{2, 7, 9, 3, 1}
	fmt.Println(rob3(arr))
}
