package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func rightSideView(root *TreeNode) []int {
	res := make([]int,0)
	dfs(root,0,&res)
	return res
}
func dfs(root *TreeNode,level int,res *[]int){
	if root == nil{
		return
	}
	if len(*res) == level{
		*res = append(*res,root.Val)
	}
	dfs(root.Left,level+1,res)
	dfs(root.Right,level+1,res)
}
func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}
	fmt.Println(rightSideView(&node1))
}
