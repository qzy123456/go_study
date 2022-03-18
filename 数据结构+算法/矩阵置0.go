package main

import "fmt"

//给一个矩阵，把0周围的都置成0
//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
func setZeroes(matrix [][]int)  {
	row := make([]bool,len(matrix))
	col := make([]bool,len(matrix[0]))
	for i, v := range matrix {
		for j, w := range v {
			if w == 0{
				row[i] = true
				col[j] = true
			}
		}
	}

	for i, v := range matrix {
		for j, _ := range v {
			if row[i] || col[j]{
				v[j] = 0
			}
		}
	}

}

func main() {
	matrix := [][]int{{1,1,1},{1,0,1},{1,1,1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
