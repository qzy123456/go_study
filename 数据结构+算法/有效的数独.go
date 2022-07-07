package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m1 := make(map[byte]bool)
		m2 := make(map[byte]bool)
		m3 := make(map[byte]bool)
		//fmt.Printf("i: %d, num[i]: %v\n", i, string(board[i]))
		// 判断每一行是否重复
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				//fmt.Printf("j1: %d, num[j]: %v\n", j, board[i][j])
				if m1[board[i][j]] {
					return false
				}
				m1[board[i][j]] = true
			}
			// 判断每一列是否重复
			if board[j][i] != '.' {
				//fmt.Printf("j2: %d, num[j]: %v\n", j, board[j][i])
				if m2[board[j][i]] {
					return false
				}
				m2[board[j][i]] = true
			}
			// 判断9宫格内的数据是否重复
			row := (i%3)*3 + j%3   //行
			col := (i/3)*3 + j/3   //列
			//fmt.Println(row,col) //00 10 20 01 11 21 02 12 22
			if board[row][col] != '.' {
				fmt.Printf("j3: %d, num[j]: %v\n", j, string(board[row][col]))
				if m3[board[row][col]] {
					return false
				}
				m3[board[row][col]] = true
			}
		}
	}
	return true
}

//官方解法
func isValidSudoku2(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	sudoku := isValidSudoku(board)
	fmt.Println(sudoku)
}
