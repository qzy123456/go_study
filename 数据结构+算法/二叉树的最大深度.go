package main

import (
	"fmt"
)

//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到叶子节点经过的最短路径上的节点数量。

type TreeNode struct {
	deep  int
	left  *TreeNode
	right *TreeNode
}
//深度优先
func maxDepth(root *TreeNode) int {
	if root == nil { //第一次进入时根节点的判断
		return 0
	}

	x := maxDepth(root.left)
	y := maxDepth(root.right)
	return max(x,y) + 1 // 1为当前节点的深度
}

func max(a,b int)int  {
	if a>b {
		return  a
	}
	return  b
}
func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}
	fmt.Println(maxDepth(&node1))
}
