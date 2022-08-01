package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTowLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTowLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTowLists(l1, l2.Next)
	return l2
}
func mergeTowListss(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	mid   := length / 2
	left  := mergeTowListss(lists[:mid])
	right := mergeTowListss(lists[mid:])
	return mergeTowLists(left, right)
}

func main() {
	var n1 = &ListNode{}
	n1.Val = 1
	n1.Next = &ListNode{Val: 2, Next: &ListNode{Val: 5,
		Next: &ListNode{Val: 6},
	}}
	var n2 = &ListNode{}
	n2.Val = 3
	n2.Next = &ListNode{Val: 4, Next: &ListNode{Val: 5,
		Next: &ListNode{Val: 6},
	}}
	var n3 = &ListNode{}
	n3.Val = 3
	n3.Next = &ListNode{Val: 4, Next: &ListNode{Val: 5,
		Next: &ListNode{Val: 6},
	}}
	//res :=  mergeTowLists(n1,n2)
	res := mergeTowListss([]*ListNode{n1, n2, n3})
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
