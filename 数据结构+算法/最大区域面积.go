package main

import "fmt"

//最大区域面积，有的也叫装水
//给出一个非负整数数组 a1，a2，a3，…… an，
//每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，选择两堵墙，和 x 轴构成的容器可以容纳最多的水。
func maxArea(arr []int) int {
	max, start, end := 0, 0, len(arr)-1
	for start < end {
		width := end - start
		high := 0
		if arr[start] < arr[end] {
			high = arr[start]
			start++
		} else {
			high = arr[end]
			end--
		}
		fmt.Println(width,high)
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println(maxArea(arr))
}
