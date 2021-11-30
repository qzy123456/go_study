package main

import "fmt"

// 解法一 原地 DP，无辅助空间
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		 //fmt.Println(grid[i][0],grid[i-1][0])  //走的时候1，1，4，2
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]

}

// 解法二 最原始的方法，辅助空间 O(n^2)
//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。说明：每次只能向下或者向右移动一步。
//在地图上求出从左上角到右下角的路径中，数字之和最小的一个，输出数字和。
//这一题最简单的想法就是用一个二维数组来 DP，当然这是最原始的做法。
// 由于只能往下和往右走，只需要维护 2 列信息就可以了，从左边推到最右边即可得到最小的解。更近一步，
// 可以直接在原来的数组中做原地 DP，空间复杂度为 0 。
func minPathSum1(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
    //初始化一个跟原数组一样大小格式的数组
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	//初始化第一列
	for i := 0; i < len(dp); i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		}
	}

	// 初始化第一行
	for i := 0; i < len(dp[0]); i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = grid[0][i] + dp[0][i-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
func min(a,b int)int  {
	if a> b {
		return b
	}
	return  a
}

func main() {
	//Input:
	//[
	//  [1,3,1],
	//  [1,5,1],
	//  [4,2,1]
	//]
	//Output: 7
	//Explanation: Because the path 1→3→1→1→1 minimizes the sum.
	arr := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	arr1 := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	fmt.Println(minPathSum(arr))
	fmt.Println(minPathSum1(arr1))
}