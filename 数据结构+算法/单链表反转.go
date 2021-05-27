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
	var next *Node = nil
	var prev *Node = nil
	var curr = head

	for curr!=nil{
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
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