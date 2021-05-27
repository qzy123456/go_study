package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ress [][]int

func levelOrder2(root *TreeNode) [][]int {
	ress = [][]int{}
	dfs(root, 0)
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

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}

func main() {
	node1 := &TreeNode{3, nil, nil}
	node1.Left = &TreeNode{9, nil, nil}
	node1.Right = &TreeNode{20, nil, nil}
	node1.Right.Left = &TreeNode{15, nil, nil}
	node1.Right.Right = &TreeNode{7, nil, nil}

	res := levelOrder(node1)
	fmt.Printf("res is: %v\n", res)
	res2 := levelOrder2(node1)
	fmt.Printf("res2 is: %v\n", res2)
}
