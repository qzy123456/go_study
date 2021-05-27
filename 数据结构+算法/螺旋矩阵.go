package main

import "fmt"

//  1--2--3
//  4--5--6
//  7--8--9
//顺时针旋转， 也就是输出 123698745
func spiralOrder(matrix [][]int) []int {
	// write code here
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var n = len(matrix[0]) //获取x轴的长度
	var m = len(matrix)    //获取y轴的长度
	var left = 0           //左边界
	var right = n - 1      //右边界
	var up = 0             //上边界
	var down = m - 1       //下边界

	var total = n * m //总的元素个数
	var res []int
	for total > 0 {
		//最上
		for i := left; i <= right && total > 0; i++ {
			//fmt.Println(matrix[up][i]) //1235
			res = append(res, matrix[up][i])
			total--
		}
		up++
		//最右
		for i := up; i <= down && total > 0; i++ {
			//fmt.Println(matrix[i][right]) //69
			res = append(res, matrix[i][right])
			total--
		}
		right--
		//最下
		for i := right; i >= left && total > 0; i-- {
			//fmt.Println(matrix[down][i]) //87
			res = append(res, matrix[down][i])
			total--
		}
		down--
		//最左
		for i := down; i >= up && total > 0; i-- {
			//fmt.Println(matrix[i][left]) //4
			res = append(res, matrix[i][left])
			total--
		}
		left++
	}
	return res
}
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}
