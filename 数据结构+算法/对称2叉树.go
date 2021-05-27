package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode)bool  {
     return isMirror(root,root)
}
func isMirror(p,q *TreeNode) bool  {
	if p==nil && q ==nil{
		return true
	}else if p==nil || q==nil{
		return false
	}
	if p.Val == q.Val{
		return isMirror(p.Left,q.Right) && isMirror(p.Right,q.Left)
	}
	return false
}

func main() {
     head := &TreeNode{
     	Val:3,
     	Left:&TreeNode{Val:1},
     	Right:&TreeNode{Val:1},
	 }
     fmt.Println(isSymmetric(head))
}
