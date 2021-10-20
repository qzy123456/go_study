package main

import "fmt"

//从 x 轴开始，给出一个数组，数组里面的数字代表从 (0,0) 点开始，宽度为 1 个单位，
// 高度为数组元素的值。如果下雨了，问这样一个容器能装多少单位的水？
func trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}

func main() {
	nums := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(nums))
}
