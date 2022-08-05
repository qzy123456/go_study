package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ress [][]int
//给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （
//即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//102题是从上到下
func levelOrderBottom(root *TreeNode) [][]int {
	ress = [][]int{}
	dfs(root, 0)
	//反转arr
	for i := 0; i < len(ress) / 2; i++ {
		ress[i], ress[len(ress) - 1 - i] = ress[len(ress) - 1 - i], ress[i]
	}
	return ress
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(ress) == level {  //这里相当于make ，不然会溢出下标ß
			fmt.Println(len(ress),level)
			ress = append(ress, []int{})
		}
		fmt.Println(ress)
		ress[level] = append(ress[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}

func main() {
	node1 := &TreeNode{3, nil, nil}
	node1.Left = &TreeNode{9, nil, nil}
	node1.Right = &TreeNode{20, nil, nil}
	node1.Right.Left = &TreeNode{15, nil, nil}
	node1.Right.Right = &TreeNode{7, nil, nil}

	res := levelOrderBottom(node1)
	fmt.Printf("res is: %v\n", res)

}
