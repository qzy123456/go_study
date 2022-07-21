package main

import (
	"fmt"
	"math"
)

//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到叶子节点经过的最短路径上的节点数量。

type TreeNode struct {
	deep  int
	left  *TreeNode
	right *TreeNode
}

//广度优先
func minDepth2(root *TreeNode) int {
	if root == nil { //第一次进入时根节点的判断
		return 0
	}

	root.deep = 1
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]  //取出第一个元素
		queue = queue[1:] // 第一个元素出队
		// 一找到叶子节点就返回
		if node.left == nil && node.right == nil {
			return node.deep
		}
		if node.left != nil {
			node.left.deep = node.deep + 1
			queue = append(queue, node.left)
		}
		if node.right != nil {
			node.right.deep = node.deep + 1
			queue = append(queue, node.right)
		}
	}

	return 0 // 错误情况
}

//深度优先
func minDepth(root *TreeNode) int {
	if root == nil { //第一次进入时根节点的判断
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}

	var min = math.MaxInt64
	if root.left != nil {
		min = mins(minDepth(root.left), min)
	}
	if root.right != nil {
		min = mins(minDepth(root.right), min)
	}
	return min + 1 // 1为当前节点的深度
}

func mins(a ,b int)int{
	if a > b{
		return b
	}
	return a
}

func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}
	fmt.Println(minDepth(&node1))
	fmt.Println(minDepth2(&node1))
}
