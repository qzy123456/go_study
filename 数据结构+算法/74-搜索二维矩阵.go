package main

import "fmt"

//给出一个二维矩阵，矩阵的特点是随着矩阵的下标增大而增大。要求设计一个算法能在这个矩阵中高效的找到一个数，
// 如果找到就输出 true，找不到就输出 false。
//虽然是一个二维矩阵，但是由于它特殊的有序性，所以完全可以按照下标把它看成一个一维矩阵，
// 只不过需要行列坐标转换。最后利用二分搜索直接搜索即可。
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, low, high := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1
	//4 0 11
	for low <= high {
		mid := low + (high-low)>>1
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1,   2,  3,  4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}

fmt.Println(searchMatrix(matrix,3))
}