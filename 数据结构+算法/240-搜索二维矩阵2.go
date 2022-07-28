package main

import "fmt"

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
// 解法一 模拟，时间复杂度 O(m+n)
func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row <= len(matrix)-1 {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			row++
		} else {
			col--
		}
	}
	return false
}

// 解法二 二分搜索，时间复杂度 O(n log n)
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for _, row := range matrix {
		low, high := 0, len(matrix[0])-1
		for low <= high {
			mid := low + (high-low)>>1
			if row[mid] > target {
				high = mid - 1
			} else if row[mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 5, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix1(matrix, 3))
	fmt.Println(searchMatrix2(matrix, 3))
}
