package main

import "fmt"

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。机器人每次只能向下或者向右移动一步。
//机器人试图达到网格的右下角（在下图中标记为“Finish”）。问总共有多少条不同的路径？

//这是一道简单的 DP 题。输出地图上从左上角走到右下角的走法数。
//由于机器人只能向右走和向下走，所以地图的第一行和第一列的走法数都是
// 1，地图中任意一点的走法数是 dp[i][j] = dp[i-1][j] + dp[i][j-1]
//1:不带障碍物
func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}

//2: 带障碍物，障碍物的处理方法是 dp[i][j]=0，需要注意的一种情况是，起点就是障碍物，那么这种情况直接输出 0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	paths := make([][]int, n+1)
	for i := range paths {
		paths[i] = make([]int, m+1)
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	paths[1][1] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if (i == 1 && j == 1) || obstacleGrid[i-1][j-1] == 1 {
				continue
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}
	return paths[n][m]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
	nums := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(nums))
}
