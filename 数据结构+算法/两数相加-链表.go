package main

import "fmt"

type ListNode struct {
      Val int
      Next *ListNode
  }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res  = &ListNode{}
	var curr  = res
	var add int

	for add!=0 || l1!=nil || l2!=nil{
		var x, y int

		if l1 != nil{
			x = l1.Val
			l1 = l1.Next
		}

		if l2!=nil{
			y = l2.Val
			l2 = l2.Next
		}

		var result = x + y + add
		add = result/10
		newNode := &ListNode{Val:result%10}
		curr.Next = newNode
		curr = newNode

	}

	return res.Next
}

func main()  {
    var n1 = &ListNode{}
    n1.Val =1
    n1.Next = &ListNode{Val:2,Next:&ListNode{Val:5}}
	var n2 = &ListNode{}
	n2.Val =9
	n2.Next = &ListNode{Val:8}
	res := addTwoNumbers(n1,n2)
	fmt.Println(res.Val,res.Next.Val,res.Next.Next.Val)
}