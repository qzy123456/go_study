package main

import "fmt"

//538. 把二叉搜索树转换为累加树
//给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），
// 使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
//
//提醒一下，二叉搜索树满足下列约束条件：
//
//节点的左子树仅包含键 小于 节点键的节点。
//节点的右子树仅包含键 大于 节点键的节点。
//左右子树也必须是二叉搜索树。

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			fmt.Println(sum)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}
//func dfs (root *TreeNode,sum *int){
//    if root != nil{
//            dfs(root.Right,sum)
//            *sum += root.Val
//            root.Val = *sum
//            dfs(root.Left,sum)
//        }
//}
func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}
	convertBST(&node1)
}
