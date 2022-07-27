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

	max1 := isRob(nums, 0, length-2)
	max2 := isRob(nums, 1, length-1)
	return max(max1, max2)
}

func isRob(nums []int, start, end int) int {
	R := 0
	nR := 0
	for i := start; i <= end; i++ {
		tempMax := max(R, nR)
		R = nums[i] + nR
		nR = tempMax
	}
	return max(R, nR)
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
