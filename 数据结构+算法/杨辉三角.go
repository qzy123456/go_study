package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	yanghui := make([][]int,numRows)
	for i := 0; i < numRows; i++ {
		// 此处需要对yanghui[i]进行make分配
		yanghui[i] = make([]int,i+1)
		for j := 0; j <= i; j++ {
			if i < 2 {
				yanghui[i][j] = 1
			} else {
				if j == 0 || j == i {
					yanghui[i][j] = 1
				} else {
					yanghui[i][j] = yanghui[i-1][j-1] + yanghui[i-1][j]
				}
			}
		}
	}
	return yanghui
}

func main() {
  fmt.Println(generate(5))
}
