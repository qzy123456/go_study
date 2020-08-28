package main

import (
	"fmt"
)

//假设是一群小孩在玩这个游戏
//创建一个小孩的结构体
type BoyNode struct {
	No int //给每个小孩一个唯一的身份编号
	next *BoyNode //指向下一个小孩
}

//假设有number个小孩在玩游戏
func AddBoyNode(number int)  *BoyNode {
	head := &BoyNode {} //先创建一个头节点
	temp := &BoyNode {} //创建一个辅助节点
	for i := 1; i<= number; i++ {
		boy := &BoyNode {
			No : i,
		}
		if i == 1 {
			head = boy
			temp = boy
			temp.next = head
		} else {
			temp.next = boy
			temp.next.next = head
		}
		temp = temp.next
	}
	return head
}

//开始游戏
func Play(head *BoyNode, a, b int) *BoyNode {
	temp := head   //辅助节点指向头节点head
	helper := head  //辅助节点指向链表最后的节点
	//判断如果没有小孩无法游戏
	if temp.next == nil {
		fmt.Println("没有小孩，无法进行游戏！")
		return head
	}
	//小孩只剩一个的时候退出游戏
	if temp.next == head {
		fmt.Println("只有一个小孩了，游戏结束！最后一个小孩为：")
		return head
	}
	//将helper指向最后一个小孩
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//循环找到从第a个小孩开始游戏
	for {
		if temp.No == a {
			break
		}
		temp = temp.next
		helper = helper.next
	}
	//数到b时出列的小孩
	for i := 1; i < b; i++ {
		temp = temp.next
		helper = helper.next
	}
	//如果该小孩是第一个小孩，则将头节点指向下一个小孩
	if temp == head {
		head = head.next
	}
	helper.next = temp.next

	fmt.Println()
	fmt.Printf("小男孩：%d出列！",temp.No)
	fmt.Println()
	ListBoyNode(head)
	//下一次游戏从第a个小孩开始
	a = temp.next.No
	return Play(head,a,b)
}

//输出显示链表
func ListBoyNode(head *BoyNode) {
	temp := head
	for {
		fmt.Printf("小男孩：%d ==>",temp.No)
		temp = temp.next
		if temp == head {
			break
		}
	}
}


func main() {
	fmt.Println("请输入有多少个小孩玩游戏：")
	var number int
	fmt.Scan(&number)
	head := AddBoyNode(number)
	ListBoyNode(head)
	fmt.Println("请输入从第几个小孩开始游戏：")
	var a int
	fmt.Scan(&a)
	fmt.Println("请输入数几的小男孩出列：")
	var b int
	fmt.Scan(&b)
	head = Play(head,a,b)
	ListBoyNode(head)
}