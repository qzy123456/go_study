package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//有序链表
func deleteDuplicat(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var current = head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			// 即1-》1-》2
			current.Next = current.Next.Next //下个元素变成下一个的下一个 1-》2
		} else {
			current = current.Next
		}
	}
	return head
}

func main() {
	var n1 = &ListNode{}
	n1.Val = 1
	n1.Next = &ListNode{Val: 1,
						Next: &ListNode{Val: 5,
										Next: &ListNode{Val: 6},
	}}
	res := deleteDuplicat(n1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
