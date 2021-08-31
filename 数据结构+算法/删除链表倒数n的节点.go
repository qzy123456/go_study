package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
//删除链表倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next:head}
	preSlow,slow,fast := dummyHead,head,head
	for fast != nil{
		if n<=0{
			//慢节点跑到末尾+n， end +n
			preSlow = slow
			slow = slow.Next
		}
		n--
		//fast跑到最后
		fast = fast.Next
	}
	//把要删除的过滤掉
	preSlow.Next = slow.Next
	return dummyHead.Next
}

func main() {
	var n1 = &ListNode{}
	n1.Val = 1
	n1.Next = &ListNode{Val: 1,
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 6},
		}}
	res := removeNthFromEnd(n1,4)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}