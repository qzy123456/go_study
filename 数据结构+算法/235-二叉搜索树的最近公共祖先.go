package main

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
//二叉搜索树，所有的左节点小于根节点，所有的右子树，大于根节点
//最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
// 最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if p.Val == q.Val {
		return p
	}
	if p.Val < root.Val && q.Val < root.Val{
		return lowestCommonAncestor1(root.Left,p,q)
	}else if p.Val > root.Val && q.Val > root.Val{
		return lowestCommonAncestor1(root.Right,p,q)
	}else{
		return root
	}
}
