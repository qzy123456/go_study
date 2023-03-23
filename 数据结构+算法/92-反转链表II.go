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
	//0->1->2->3->4->5
	//fmt.Println(pre.Value,m_node.Value,n_node.Value) //0.1.4
	for m_node != n_node{
		fmt.Println("1=>",m_node.Next.Value)
		pre.Next = m_node.Next
		m_node.Next = n_node.Next
		fmt.Println("2=>",n_node.Next.Value)
		n_node.Next = m_node
		fmt.Println("3=>",m_node.Value)
		m_node = pre.Next
		fmt.Println("4=>",m_node.Value)
		fmt.Println()
	}
   return dummy.Next
}

/***
  头插法
 */
func reverseBetween2(head *Node, left, right int) *Node {
	// 设置 dummyNode 是这一类问题的一般做法，防止第一个节点就开始反转
	dummyNode := &Node{}
	dummyNode.Next = head
	pre := dummyNode
	//pre移到开启的之前
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	//fmt.Println(pre.Value) //1
	//cur就是要开始逸动的节点
	cur := pre.Next
	//fmt.Println(cur.Value)  //2
	for i := 0; i < right-left; i++ {
		next := cur.Next
		//fmt.Println(next.Value) //3
		cur.Next = next.Next
		//fmt.Println(cur.Next.Value) //4
		next.Next = pre.Next
		//fmt.Println(next.Next.Value) //2
		pre.Next = next
		//fmt.Println(pre.Next.Value)//3
	}
	//第一轮 // 1  -3 - 2 - 4 -5
	//         pre     cur
	//第2轮 // 1-  4 - 3 -  2   -5
	//        pre          cur
	//第3轮 // 1   -5-4-3- 2
	//        pre	      cur
	return dummyNode.Next
}

func main() {
	var node5 = Node{5, nil}
	var node4 = Node{4, &node5}
	var node3 = Node{3, &node4}
	var node2 = Node{2, &node3}
	var node1 = Node{1, &node2}
	res := reverseBetween2(&node1,2,5)
	for res !=nil{
		fmt.Print("value:", res.Value, " ->")
		res = res.Next
	}
}
