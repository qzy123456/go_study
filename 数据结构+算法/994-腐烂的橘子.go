package main

import "fmt"

//994. 腐烂的橘子
//在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
//
//值 0 代表空单元格；
//值 1 代表新鲜橘子；
//值 2 代表腐烂的橘子。
//每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
//返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1
func orangesRotting(grid [][]int) int {
	//如果没有新鲜的，不用处理
	if !hasFresh(grid) {
		return 0
	}

	time := getTimes(grid)
    //感染完，发现还有没有被感染的，返回-1
	if hasFresh(grid) {
		return -1
	}
	return time
}

func hasFresh(grid [][]int) bool {
	for i := range grid {
		for _, val := range grid[i] {
			if 1 == val {
				return true
			}
		}
	}
	return false
}

func getTimes(grid [][]int) int {
	//获取腐烂的橙子
	q := getQueue(grid)
	time := 0

	for len(q) > 0 {
		next := make([][]int, 0)

		for i := range q {
			row := q[i][0]
			col := q[i][1]

			next = setRot(grid, next, row-1, col)
			next = setRot(grid, next, row, col+1)
			next = setRot(grid, next, row+1, col)
			next = setRot(grid, next, row, col-1)
		}

		if len(next) != 0 {
			time++
		}
		q = next
	}
	return time
}

func getQueue(grid [][]int) [][]int {
	q := make([][]int, 0)
	for i := range grid {
		for j, val := range grid[i] {
			if 2 == val {
				q = append(q, []int{i, j})
			}
		}
	}
	return q
}

func setRot(grid, q [][]int, row, col int) [][]int {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] != 1 {
		return q
	}
	grid[row][col] = 2
	q = append(q, []int{row, col})
	return q
}

func main() {
	grid := [][]int{{2,1,1},{1,1,0},{0,1,1}}
	fmt.Println(orangesRotting(grid))
}
