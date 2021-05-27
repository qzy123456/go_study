package main

import (
	"fmt"
)

//给定一个链表，返回链表开始入环的第一个节点，如果链表无环，则返回null

type ListNode struct {
	Val int
	Next *ListNode
}

type void struct {}
//1：直接遍历链表，使用一个值作为标记位（标记是否已经到达过）
func detecCycle(head *ListNode)*ListNode  {
	if head ==nil || head.Next ==nil{
		return nil
	}
	var nodes  = map[*ListNode]void{}
	for head != nil {
		if _,ok := nodes[head];ok{
			return head
		}
		nodes[head] = void{}
		head = head.Next
	}
	return nil
}
//2快慢指针
func detecCycle2(head *ListNode)*ListNode  {
	if head==nil||head.Next==nil{
		return nil
	}

	var flag = false//是否存在环
	var slow = head
	var quick = head

	for quick!=nil && quick.Next!=nil{
		slow = slow.Next
		quick = quick.Next.Next
		if quick==slow{
			flag = true
			break
		}
	}

	if flag{
		fmt.Println(quick == head)
		fmt.Println(slow)
		//false
		//&{5 0xc00008e070}
		slow = head
		for slow!=quick{
			slow = slow.Next
			quick = quick.Next
		}
		return slow
	}

	return nil

}

func main() {
	var head = &ListNode{}
	  node1 := &ListNode{Val:3}
	  node2 := &ListNode{Val:2}
	  node3 := &ListNode{Val:5}
	  node4 := &ListNode{Val:4}
      head.Next=node1
      node1.Next=node2
      node2.Next=node3
      node3.Next=node4
      node4.Next=node2
      fmt.Println(detecCycle(head))
      fmt.Println(detecCycle2(head))
}