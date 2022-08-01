package main

import "fmt"

func rotates(matrix [][]int)  {
    //左右交换
    length := len(matrix)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

    //中间截取，上下交换
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-1-j] = matrix[i][length-1-j], matrix[i][j]
		}
	}
}

func main() {
     var nums = [][]int{
     	{1,2,3},
     	{4,5,6},
     	{7,8,9},
     }
     rotates(nums)
     fmt.Println(nums)
}
