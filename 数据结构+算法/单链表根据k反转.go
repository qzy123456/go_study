package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//Given this linked list: 1->2->3->4
//For k = 2, you should return: 2->1->4->3
//For k = 3, you should return: 3->2->1->4
//按照每 K 个元素翻转的方式翻转链表。如果不满足 K 个元素的就不翻转
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	//取第k+1，node = k+ 1
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	//fmt.Printf("%+v\n",node) //4
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	//fmt.Printf("%+v\n",first) //1
	//fmt.Printf("%#v\n",last)  //4
	for first != last {
		//fmt.Printf("%#v\n",first) // 1 2 3
		tmp := first.Next  //2
		first.Next = prev  //4
		prev = first    //1
		first = tmp     //2
	}
	return prev
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	dummy, curr, tmp, len := &ListNode{Next: head}, head, &ListNode{Next: nil}, 0
	prev := dummy
	for head != nil {
		len++
		head = head.Next
	}
	head = dummy.Next
	for i := 0; i < len/k; i++ {
		for j := 0; j < k-1; j++ {
			// 123   213   321
			tmp = curr.Next // 2
			curr.Next = tmp.Next
			tmp.Next = prev.Next // 213
			prev.Next = tmp      //321
		}
		prev = curr
		curr = prev.Next
	}
	return dummy.Next
}

func main() {
	var n1 = &ListNode{}
	n1.Val = 1
	n1.Next = &ListNode{
					Val: 2,
					Next: &ListNode{
								Val: 3,
								Next: &ListNode{
											Val: 4},
	}}
	res := reverseKGroup(n1,3)
	//res := reverseKGroup2(n1, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
