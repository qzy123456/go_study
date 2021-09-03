package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
//删除链表倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	// 快指针向后移N步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 如果快指针移到了最后，说明被删除的结点应该是链表的头结点
	if fast == nil {
		head = head.Next
		return head
	}
	// 将快指针和慢指针同时向后移动，直到快指针移到了链表的最后一个结点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 删除结点
	slow.Next = slow.Next.Next
	return head
}

func main() {
	var n1 = &ListNode{}
	n1.Val = 1
	n1.Next = &ListNode{Val: 2,
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 6},
		}}
	res := removeNthFromEnd(n1,2)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}
}