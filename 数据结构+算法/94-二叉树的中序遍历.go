package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}
func inorderTraversal1(root *TreeNode) []int {
	var result []int
	pre(root, &result)
	return result
}
func inorderTraversal2(root *TreeNode) []int {
	var result []int
	after(root, &result)
	return result
}

//中序
func inorder(root *TreeNode, output *[]int) {
	//中序遍历:先遍历左子树再遍历根，再遍历根节点再遍历右子树
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}
//前序pre
func pre(root *TreeNode, output *[]int) {
	//前序遍历：先遍历根节点再遍历左子树，再遍历右子树
	if root != nil {
		*output = append(*output, root.Val)
		pre(root.Left, output)
		pre(root.Right, output)
	}
}
//后序
func after(root *TreeNode, output *[]int) {

	//后序遍历:先遍历左子树再遍历根，再遍历右子树，再遍历根节点
	if root != nil {
		after(root.Left, output)
		after(root.Right, output)
		*output = append(*output, root.Val)
	}
}

func main() {
     var root = &TreeNode{
     	Val:1,
     	Left:&TreeNode{
     		Val:2,
		},
		Right:&TreeNode{
			Val:3,
		},
     }
     fmt.Println(inorderTraversal(root)) //中 213
     fmt.Println(inorderTraversal1(root)) //前 123
     fmt.Println(inorderTraversal2(root)) //后 231
}
