package main

import "fmt"

type Node struct{
	Value int
	Next *Node
}
func reverseBetween(head *Node , m int,n int) *Node {
   if m == n{
   	return head
   }
   var dummy = &Node{}
   dummy.Next = head
   pre := dummy
   m_node := head
   n_node := head
	for i:=1;i<m;i++ {
		pre = pre.Next
		m_node = m_node.Next
	}
	for i:=1;i<n;i++ {
		n_node = n_node.Next
	}
	for m_node != n_node{
		pre.Next = m_node.Next
		m_node.Next = n_node.Next
		n_node.Next = m_node
		m_node = pre.Next
	}
   return dummy.Next
}

func main() {
	var node5 = Node{5, nil}
	var node4 = Node{4, &node5}
	var node3 = Node{3, &node4}
	var node2 = Node{2, &node3}
	var node1 = Node{1, &node2}
	res := reverseBetween(&node1,2,4)
	for res !=nil{
		fmt.Print("value:", res.Value, " ->")
		res = res.Next
	}
}
