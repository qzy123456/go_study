package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
var sum int
var L int
var R int

func rangeSumBST(root *TreeNode,l,r int) int {
	L = l
	R = r
	travle(root)
	return sum
}
func travle(root *TreeNode)  {
	if root == nil{
		return
	}
	addValue(root.Val)
	travle(root.Left)
	travle(root.Right)
}
func addValue(val int)  {
	if val >= L && val <= R{
		sum+=val
	}
}
func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}

	fmt.Println(rangeSumBST(&node1,2,3))
}
