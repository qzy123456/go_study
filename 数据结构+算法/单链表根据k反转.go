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
	//fmt.Printf("%+v\n",node) //3
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	fmt.Println("刚进来",first,last)  //1,3
	prev := last
	for first != last{
		temp := first.Next
		fmt.Println("temp",temp) //2，3
		first.Next = prev
		prev = first
		fmt.Println("prev",prev)//1 ，2
		first = temp
		fmt.Println("first",first)//2 ，3
	}
	return prev
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	length:=0
	pre:=head
	//计算链表长度
	for pre!=nil{
		length++
		pre = pre.Next
	}

	//k个一组的次数
	time:=length/k
	//临时节点
	dummy:=new(ListNode)
	dummy.Next = head
	pre = dummy    //每一组前面的数，头插法，每次都把move放在pre后面
	var nexthead,move *ListNode

	for i:=0;i<time;i++{
		nexthead = pre.Next    //每个k长度的开头，逐渐变为末尾
		move = nexthead.Next//第二个，不断向后移动，把move插入到pre后面，
		for j:=0;j<k-1;j++{
			nexthead.Next = move.Next
			move.Next = pre.Next
			pre.Next = move
			move = nexthead.Next
		}
		//经过头插法，nexthead逐渐变为结尾，结束一轮循环时，nexthead变为结尾，是下一个k长度的pre，
		pre = nexthead  //下一个k同样的方式，头插法放pre后面

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
	res := reverseKGroup(n1,2)
	//res := reverseKGroup2(n1, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
