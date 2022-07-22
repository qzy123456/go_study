package main

import "fmt"

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

//给出 2 个数组，根据 preorder 和 inorder 数组构造一颗树。
//利用递归思想，从 preorder 可以得到根节点，
// 从 inorder 中得到左子树和右子树。只剩一个节点的时候即为根节点。不断的递归直到所有的树都生成完成。
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for index, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(preorder[1:index+1], inorder[:index])
			root.Right = buildTree(preorder[index+1:], inorder[index+1:])
		}
	}
	return root
}


func main()  {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	fmt.Println(buildTree(preorder,inorder))
}
