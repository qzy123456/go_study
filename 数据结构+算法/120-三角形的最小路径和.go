package main

import (
	"fmt"
	"math"
)

//120. 三角形最小路径和
//给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
//每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
//
//
//示例 1：
//
//输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//输出：11
//解释：如下面简图所示：
//   2
//  3 4
// 6 5 7
//4 1 8 3
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//示例 2：
//
//输入：triangle = [[-10]]
//输出：-10
func minimumTotal(triangle [][]int) int {
	var minV = math.MaxInt64
	for i := range triangle {
		for j := range triangle[i] {
			if i == 0 && j == 0 {
				continue
			}
			if j == 0 {
				// j=0的情况，取上面+当前
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				// j=最后一个，取斜上+当前
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				// 其他，取斜上、上面的最小+当前
				triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}
	// 计算最后一行的最小值
	for _, v := range triangle[len(triangle)-1] {
		if v < minV {
			minV = v
		}
	}
	return minV
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	triangle := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
     fmt.Println(minimumTotal(triangle))
}
