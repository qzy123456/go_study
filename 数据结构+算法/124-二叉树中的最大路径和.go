package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给出一个二叉树，要求找一条路径使得路径的和是最大的。
//这一题思路比较简单，递归维护最大值即可。不过需要比较的对象比较多。maxPathSum(root) = max(maxPathSum(root.Left), maxPathSum(root.Right), maxPathSumFrom(root.Left) (if>0) + maxPathSumFrom(root.Right) (if>0) + root.Val) ，其中，maxPathSumFrom(root) = max(maxPathSumFrom(root.Left), maxPathSumFrom(root.Right)) + root.Val
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		//左+根+右 情况
		all := left + root.Val + right
		//更新到全局变量
		maxSum = max(maxSum, all)
		//max(左，右)+根 --->单边 情况
		half := root.Val + max(left, right)
		//主要是为了应为负数的情况，既然是负数，还不如不加
		return max(half, 0)
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//124.二叉树中的最大路径和
func maxPathSum2(root *TreeNode) int {
	var maxSum = make([]int,1)
	maxSum[0] = -1 << 63
	maxPathSumHelper(root, maxSum)
	return maxSum[0]
}

//124.二叉树中的最大路径和-辅助函数
func maxPathSumHelper(root *TreeNode, maxSum []int) int {
	if root == nil {
		return 0
	}
	lSum := max(maxPathSumHelper(root.Left, maxSum), 0) //保留左边的最大值，最小为0
	rSum := max(maxPathSumHelper(root.Right, maxSum), 0) //右边最大值
	aSum := root.Val + lSum + rSum    //临时和
	maxSum[0] = max(aSum, maxSum[0])  //最大值
	return root.Val + max(rSum, lSum)
}

func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}
	fmt.Println(maxPathSum2(&node1))
	fmt.Println(maxPathSum(&node1))
}
