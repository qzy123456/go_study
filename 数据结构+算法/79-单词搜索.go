package main

import "fmt"

func exist(board [][]byte, word string) bool {
	//要搜索的单词不存在
	if word == "" {
		return false
	}

	//x轴
	m := len(board)
	if m == 0 {
		return false
	}
    //y轴
	n := len(board[0])
	if n == 0 {
		return false
	}

	var dfs func(int, int, int) bool

	dfs = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}
        //上下左右四个峰值不能回退了。
		if r < 0 || r >= m || c < 0 || c >= n || word[index] != board[r][c] {
			return false
		}
        //标志已经回溯过了
		temp := board[r][c]
		board[r][c] = 0
		//上下左右四个方向，index是单词的下标
		if dfs(r+1, c, index+1) || dfs(r, c+1, index+1) || dfs(r-1, c, index+1) || dfs(r, c-1, index+1) {
			return true
		}
		//重新回溯
		board[r][c] = temp

		return false
	}

	for i := range board {
		for j := range board[0] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}


func main() {
	board  := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word := "ABCCED"
	fmt.Println(exist(board,word))
}

