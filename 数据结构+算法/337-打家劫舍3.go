package main

import "fmt"

//小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
//如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//示例 1:
//输入: root = [3,2,3,null,3,null,1]
//    3
//    / \
//   2   3
//    \   \
//     3   1
//
//Output: 7
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//解题思路 #
//这一题是打家劫舍的第 3 题。这一题需要偷的房子是树状的。报警的条件还是相邻的房子如果都被偷了，就会触发报警。
//只不过这里相邻的房子是树上的。问小偷在不触发报警的条件下最终能偷的最高金额。
//解题思路是 DFS。当前节点是否被打劫，会产生 2 种结果。如果当前节点被打劫，那么它的孩子节点可以被打劫；
//如果当前节点没有被打劫，那么它的孩子节点不能被打劫。按照这个逻辑递归，最终递归到根节点，取最大值输出即可。
func rob(root *TreeNode) int {
	a, b := dfsTreeRob(root)
	return max(a, b)
}

func dfsTreeRob(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}
	l0, l1 := dfsTreeRob(root.Left)
	r0, r1 := dfsTreeRob(root.Right)
	// 当前节点没有被打劫
	tmp0 := max(l0, l1) + max(r0, r1)
	// 当前节点被打劫
	tmp1 := root.Val + l0 + r0
	return tmp0, tmp1
}
func max(a,b int)int{
	if a>b{
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

	fmt.Println(rob(&node1))
}
