package main

import (
	"fmt"
)

type Node struct{
	Value int
	Next *Node
}

func PrintNodeList(head *Node){
	fmt.Print("value:", head.Value, " ->")
	for head.Next!=nil{
		head = head.Next
		fmt.Print("value:", head.Value, " ->")
	}
	fmt.Println()
}

func ReverseNodeList1(head *Node)*Node{
	//定义一个空节点，用于返回
	var resHead *Node
	//临时节点，用于循环
	now := head
	for now != nil {
		//临时节点，防止头节点丢失，用于重新指向
		temp := now
		//头节点指向下一个
		now = now.Next
		//相当于断链 1-nil ，2-1-nil ，3-2-1-nil
		temp.Next = resHead
		//把新节点拼到一起返回
		resHead = temp
	}

	return resHead
}
//反转链表的实现
//那么这道题其实就是把指针指向前一个节点
//
//位置调换次数	pre	  			cur	        	whole
//0	            nil				1->2->3->4->5	1->2->3->4->5
//1	            1->nil			2->-3>->4->5	2->3->4->5->1->nil
//2	            2->1->nil		3->4->5			3->4->5->2->1->nil
//3	 			3->2->1->nil	4->5			4->5->3->2->1->nil
//4				4->3->2->1->nil	5				5->4->3->2->1->nil
//可以看出来
//
//pre是cur的最前面那位（pre = cur）
//cur就是当前位的后面链表元素（cur = cur.Next）
//cur.Next肯定是接pre（cur.Next = pre）
func reversrList(head *Node) *Node {
	cur := head
	var pre *Node = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}

func ReverseNodeList2(head *Node)*Node{
	if head==nil || head.Next==nil{
		return head
	}

	var newHead = ReverseNodeList2(head.Next)
	head.Next.Next = head
	head.Next = nil//避免成环
	return newHead
}

func main(){
	var node5 = Node{5, nil}
	var node4 = Node{4, &node5}
	var node3 = Node{3, &node4}
	var node2 = Node{2, &node3}
	var node1 = Node{1, &node2}
	PrintNodeList(&node1)
	PrintNodeList(ReverseNodeList1(&node1))
	PrintNodeList(ReverseNodeList2(&node5))
}