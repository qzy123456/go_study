package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func binaryTreePaths(root *TreeNode) []string {
	// 如果根节点为空，则返回空节点
	if root == nil {
		return []string{}
	}

	// 如果根节点没有子节点，则直接将该点的值返回出去
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	// 创建结果字符串列表
	res := []string{}

	// 搜寻左子节点
	tmpLeft := binaryTreePaths(root.Left)
	for i := 0; i < len(tmpLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
	}

	// 搜寻右子节点
	tmpRight := binaryTreePaths(root.Right)
	for i := 0; i < len(tmpRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpRight[i])
	}
	return res
}

func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}
	fmt.Println(binaryTreePaths(&node1))
}
