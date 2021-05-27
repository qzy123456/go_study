package main

import "fmt"

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func sortArrayToBST(nums []int) *TreeNode1  {
	l := len(nums)
	if l ==0 {
		return nil
	}
	n := (l-1) /2
	node := &TreeNode1{}
	node.Val = nums[n]
	node.Left = sortArrayToBST(nums[:n])
	node.Right = sortArrayToBST(nums[n+1:])
	return node
}

func main() {
	nums := []int{1,2,3,4,5,6}
	res2 := sortArrayToBST(nums)
	result := fmt.Sprintf("%+v",*res2)
	fmt.Print(result)
}
