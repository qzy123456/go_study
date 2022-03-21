package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTowLists(l1 ,l2 *ListNode)*ListNode  {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	if l1.Val < l2.Val{
		l1.Next =mergeTowLists(l1.Next,l2)
		return l1
	}
	l2.Next = mergeTowLists(l1,l2.Next)
	return l2
}
func mergeTowLists2(l1 ,l2 *ListNode)*ListNode  {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	//设置一个虚拟头节点
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 !=nil{
		if l1.Val >= l2.Val{
			cur.Next = l2
			l2 = l2.Next
		}else {
			cur.Next = l1
			l1=l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil{
		cur.Next = l1
	}
	if l2 != nil{
		cur.Next = l2
	}
	return head.Next
}

func main() {
	var n1 = &ListNode{}
	n1.Val =1
	n1.Next = &ListNode{Val:2,Next:&ListNode{Val:5,
		Next:&ListNode{Val:6},
	}}
	var n2 = &ListNode{}
	n2.Val =3
	n2.Next = &ListNode{Val:4,Next:&ListNode{Val:5,
		Next:&ListNode{Val:6},
	}}
	//res :=  mergeTowLists(n1,n2)
	res :=  mergeTowLists2(n1,n2)
	for res !=nil{
		fmt.Println(res.Val)
		res = res.Next
	}
}


