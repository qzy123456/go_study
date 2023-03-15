package main

import (
	"fmt"
	"net"
)

// 定义节点
type Node struct {
	Value int
	Next  *Node
}

// 初始化头节点
var head = new(Node)

// 添加节点
func addNode(t *Node, v int) {
	if head == nil {
		t = &Node{v, nil}
		head = t
		return
	}

	if v == t.Value {
		fmt.Println("节点已存在:", v)
		return
	}

	// 如果当前节点下一个节点为空
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return
	}

	// 如果当前节点下一个节点不为空
	addNode(t.Next, v)
}

// 遍历链表
func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

// 查找节点
func lookupNode(t *Node, v int) bool {
	if head == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

// 获取链表长度
func size(t *Node) int {
	if t == nil {
		fmt.Println("-> 空链表!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}

// 入口函数
func main() {
	fmt.Println(head)
	head = nil
	// 添加节点
	addNode(head, 1)
	addNode(head, 2)
	// 遍历
	traverse(head)
	// 添加更多节点
	addNode(head, 10)
	addNode(head, 5)
	addNode(head, 45)
	// 添加已存在节点
	addNode(head, 5)
	// 再次遍历
	traverse(head)

	// 查找已存在节点
	if lookupNode(head, 5) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}

	// 查找不存在节点
	if lookupNode(head, -100) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}
	fmt.Println("连表长度!", size(head))
}
