package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapParis(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}
func swapPairs1(head *ListNode) *ListNode {
	//如果为空或为单数节点   直接返回
	if head == nil || head.Next == nil {
		return head
	}
	//记录第一二个节点 对第三个节点开头的子链递归 记录其返回头为newHead
	first := head
	second := head.Next
	newHead := swapPairs1(second.Next)
	//将第一二个节点排序完成  将排好序的子链拼接上去 返回新的头
	second.Next = first
	first.Next = newHead
	return second
}

func swapPairs4(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

func swapPairs3(head *ListNode) *ListNode {
	if (head == nil) || (head.Next == nil) {
		return head
	}
	nextNode := head.Next
	head.Next = swapPairs3(nextNode.Next)
	nextNode.Next = head
	return nextNode
}

func main() {
	var n1 = &ListNode{}
	n1.Val = 1
	n1.Next = &ListNode{Val: 2, Next: &ListNode{Val: 5,
		Next: &ListNode{Val: 6},
	}}
	//res := swapParis(n1)
	res := swapPairs3(n1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
