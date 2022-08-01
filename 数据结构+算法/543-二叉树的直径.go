package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//首先我们知道一条路径的长度为该路径经过的节点数减一，所以求直径（即求路径长度的最大值）等效于求路径经过节点数的最大值减一。
//
//而任意一条路径均可以被看作由某个节点为起点，从其左儿子和右儿子向下遍历的路径拼接得到。
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	depth(root, &ans)
	return ans - 1
}
func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left, ans)
	r := depth(root.Right, ans)
	*ans = max(*ans, l+r+1)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	var node7 = TreeNode{7, nil, nil}
	var node6 = TreeNode{6, &node7, nil}
	var node5 = TreeNode{5, nil, nil}
	var node4 = TreeNode{4, nil, nil}
	var node3 = TreeNode{3, &node6, nil}
	var node2 = TreeNode{2, &node4, &node5}
	var node1 = TreeNode{7, &node2, &node3}
	diameterOfBinaryTree(&node1)
}
