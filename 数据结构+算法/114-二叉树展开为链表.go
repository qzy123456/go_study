package main

import (
	_"fmt"
)

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }


func flatten(root *TreeNode)  {
	list := preorderTraversal(root)
	cur := root
	for i := 1; i < len(list); i++ {
		cur.Left = nil
		cur.Right = &TreeNode{Val: list[i].Val, Left: nil, Right: nil}
		cur = cur.Right
	}
	return
}
//前序遍历
func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	left  := root.Left
	right := root.Right
	cur := left
	if left != nil {
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = right
		root.Left = nil
		root.Right = left
	}
	flatten(root.Right)
}

func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}
	flatten(&node1)
}
