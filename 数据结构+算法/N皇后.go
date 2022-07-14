package main

import "fmt"

var res [][]string
//竖轴，撇，捺
var columns,pie,na = map[int]bool{},map[int]bool{},map[int]bool{}
//棋盘
var queens = map[int]map[int]bool{}//定义一个二维数组的类型
var queen int
func solveNQueens(n int) [][]string {
	res = [][]string{}
	//棋盘初始化
	for i:=0;i<n;i++{
		queens[i] = make(map[int]bool)
	}
	queen = n
	backtrack(0)
	return res
}

func backtrack(row int) {
	//退出条件，已经跑到最后一行
	if row == queen {
		generateBoard(queen)
		return
	}
	for i := 0; i < queen; i++ {
		if columns[i] || pie[row - i] || na[row + i]{
			continue
		}

		queens[row][i] = true
		columns[i] = true
		pie[row-i] = true
		na[row+i] = true
		backtrack(row + 1)
		queens[row][i] = false
		delete(columns, i)
		delete(pie, row-i)
		delete(na, row+i)
	}
}

func generateBoard( n int)  {
	board := []string{}
	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < n; j++ {
			if queens[i][j]{
				row+="Q"
			}else{
				row+="."
			}
		}
		board = append(board, row)
	}
	res = append(res, board)
}
func main(){
	fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(3))
	fmt.Println(solveNQueens(2))
	fmt.Println(solveNQueens(1))
}
