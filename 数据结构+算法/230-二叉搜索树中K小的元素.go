package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func kthSmallest(root *TreeNode, k int) int{
	leftCount := count(root)
	if leftCount + 1== k{
		return root.Val
	}else if k<=leftCount{
		return kthSmallest(root.Left,k)
	}else{
		return kthSmallest(root.Right,k-leftCount)
	}
}
func count(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return count(root.Left) + count(root.Left) + 1
}
func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}

	fmt.Println(kthSmallest(&node1,2))
}