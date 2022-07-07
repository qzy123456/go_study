package main

import (
	"encoding/binary"
	"fmt"
)

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
	nums := []int{1,2,3,4,5}
	fmt.Println(searchFirstEqualElement(nums,3))
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func isValidSudoku(board [][]byte) bool {
	 m1,m2,m3 := make(map[byte]bool),make(map[byte]bool),make(map[byte]bool)
	for i:=0;i<9 ;i++  {
		for j:=0;j<9 ;j++  {
			if board[i][j] != '.'{
				if m1[board[i][j]]{
					return false
				}
				m1[board[i][j]] = true
			}
			if board[j][i] != '.'{
				if m2[board[j][i]]{
					return false
				}
				m2[board[j][i]] = true
			}
			row := (i%3)*3 + j%3
			col := (i/3)*3 + j/3
			if board[row][col] != '.'{
				if m3[board[row][col]]{
					return false
				}
				m3[board[row][col]] = true
			}
		}
	}

	return true
}

