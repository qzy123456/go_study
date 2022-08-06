package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//先遍历A链表，记录所有节点，再遍历B链表，如果已经存在则返回该节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil{
		return nil
	}
	var nodes = make(map[*ListNode]struct{})
	for headA !=nil  {
		if _,ok := nodes[headA];!ok{
			nodes[headA]= struct{}{}
		}
		headA = headA.Next
	}
	for headB !=nil  {
		if _,ok := nodes[headB];ok{
			return headB
		}
		headB = headB.Next
	}
	return nil
}
//方法2
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA==nil||headB==nil{
		return nil
	}

	var lenA = 0
	var lenB = 0
	// 分别计算链表长度
	for p := headA; p!=nil; p=p.Next{
		lenA++
	}
	for p := headB; p!=nil; p=p.Next{
		lenB++
	}

	// 将链表头移动到同一起点
	if lenA>=lenB{
		for i:=0;i<lenA-lenB;i++{
			headA = headA.Next
		}
	}else{
		for i:=0;i<lenB-lenA;i++{
			headB = headB.Next
		}
	}

	// 同时遍历，相等则相交
	for headA!=nil{
		if headA==headB{
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
//链表A走完，跳到链表b，反之一样。数学公式 a+b = b+ a
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
func main() {
	var n1 = &ListNode{}
	no1 := &ListNode{
		Val:2,
	}
	n1.Val =1
	n1.Next = &ListNode{Val:2,Next:no1}
	var n2 = &ListNode{}
	n2.Val =9
	n2.Next = no1

	fmt.Println(getIntersectionNode(n1,n2))
	fmt.Println(getIntersectionNode2(n1,n2))
	fmt.Println(getIntersectionNode3(n1,n2))
}
