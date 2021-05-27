package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//旋转链表  head =[1,2,3,4,5] k = 2  ===>[4,5,1,2,3]
func rotateRight(head *ListNode, k int) *ListNode {
	if head==nil || head.Next==nil || k==0{
		return head
	}

	var length = 1//链表总长度
	var curr = head
	//求出链表的总长度
	for curr.Next!=nil{
		length++
		curr = curr.Next
	}
	// 刚好旋转回去
	k = k%length
	if k == 0{
		return head
	}

	// 首先首尾相连成环,找到倒数第k+1个节点，成为新的头节点
	curr.Next = head
	for i:=0;i<length-k;i++{
		curr = curr.Next
	}

	// curr此时指向新的尾节点
	newHead := curr.Next//新的尾节点的下一个节点即为新的头节点
	curr.Next = nil//断尾
	return newHead

}

func main() {
	var n1 = &ListNode{}
	n1.Val =1
	n1.Next = &ListNode{Val:2,Next:&ListNode{Val:5,
		Next:&ListNode{Val:6},
		}}
	res :=  rotateRight(n1,2)
	for res !=nil{
		fmt.Println(res.Val)
		res = res.Next
	}
}