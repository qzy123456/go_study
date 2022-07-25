package main

import "fmt"

//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//1：暴力
func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide
	}
	rows, columns := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide, 1)
				curMaxSide := min(rows - i, columns - j)
				for k := 1; k < curMaxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for m := 0; m < k; m++ {
						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k + 1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//动态规划
func maximalSquare2(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0{
		return 0
	}
	cols := len(matrix[0])
	maxSideLen := 0

	dp := make([][]int, rows + 1)
	dp[0] = make([]int, cols + 1)

	for i := 1; i <= rows; i++{
		dp[i] = make([]int, cols + 1)
		for j := 1; j <= cols; j++{
			if matrix[i-1][j-1] == '0'{
				dp[i][j] = 0
			}else{
				dp[i][j] = MinThree(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > maxSideLen{
					maxSideLen = dp[i][j]
				}
			}
		}
	}

	return maxSideLen*maxSideLen
}

func MinThree(a, b, c int) int {
	var min int
	if a >= b {
		min = b
	} else {
		min = a
	}

	if min >= c{
		return c
	}else{
		return min
	}
}

func main() {
	matrix := [][]byte{
		{'1','0','1','0','0'},
		{'1','0','1','1','1'},
		{'1','1','1','1','1'},
		{'1','0','0','1','0'},
	}
	fmt.Println(maximalSquare2(matrix))
	fmt.Println(maximalSquare(matrix))

}
