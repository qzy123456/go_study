package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
//有效 二叉搜索树定义如下：
//节点的左子树只包含 小于 当前节点的数。
//节点的右子树只包含 大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
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
	fmt.Println(isValidBST(root))
}
