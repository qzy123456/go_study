package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

//直接遍历 存储标记位
func hasCycle(head *Node) bool {
	st := map[Node]struct{}{}
	for head != nil {
		if _, ok := st[*head]; ok {
			return true
		}
		st[*head] = struct{}{}
		head = head.next
	}
	return false
}

//双指针
func hasCycle2(head *Node) bool {
	var slow = head
	var quick = head
	for quick != nil && quick.next != nil {
		slow = slow.next
		quick = quick.next.next
		if quick == slow {
			return true
		}
	}

	return false
}

func main() {
	node5 := Node{5, nil}
	node4 := Node{4, &node5}
	node3 := Node{3, &node4}
	node2 := Node{2, &node3}
	node1 := Node{1, &node2}
	node5.next = &node3
	fmt.Println(hasCycle(&node1))
	fmt.Println(hasCycle2(&node1))
}
