package main

import (
	"fmt"
)

//环形链表测试用结构体
type TestNode struct {
	no int //编号
	name string //姓名
	next *TestNode
}

//环形链表插入
func InsertNod(head *TestNode, newNode *TestNode) {

	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.next = head
}

//删除
func DelNode(head *TestNode, id int ) *TestNode {
	temp := head  //辅助节点指向链表头节点head
	helper := head  //辅助节点指向链表最后一个节点

	//判断链表为空
	if head.next == nil {
		fmt.Println("此链表为空！")
		return head
	}
	//判断只有一个节点，且这个节点的no编号等于id时删除
	if head.next == head && head.no == id {
		head.next = nil
		return head
	}

	//将helper指向链表最后一个节点
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		if temp.next == head {
			break
		}
		if temp.no == id {
			if temp == head {
				head = head.next
			}
			helper.next = temp.next
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}

	if flag {
		if temp.no == id {
			helper.next = temp.next
		} else {
			fmt.Printf("无id为%d的值!",id)
		}
	}
	return head

}

//显示链表
func ListNode(head *TestNode) {
	temp := head
	if temp.next == nil {
		return
	}

	for {
		fmt.Printf("%d:%s\t",temp.no,temp.name)
		temp = temp.next
		if temp == head {
			break
		}
	}
}

func main() {

	head := &TestNode{}

	t1 := &TestNode {
		no : 1,
		name : "a",
	}

	t2 := &TestNode {
		no : 2,
		name : "b",
	}

	t3 := &TestNode {
		no : 3,
		name : "c",
	}
	InsertNod(head,t1)
	InsertNod(head,t2)
	InsertNod(head,t3)
	ListNode(head)
	fmt.Println()
	 head = DelNode(head,1)
	ListNode(head)
}