package main

import (
	"fmt"
	"strconv"
)

// 定义员工结构体
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 显示员工信息
func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员id %d 名字%s\n", this.Id%7, this.Id,this.Name)
}

// EmpLink
// 这里的 EmpLink 不带表头，即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

// 1 添加员工的方法, 保证添加时，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   // 辅助指针
	var pre *Emp = nil // 辅助指针 pre 在 cur 前面
	// 如果当前的 EmpLink 就是一个空链表
	if cur == nil {
		this.Head = emp // 完成
		return
	}
	// 如果不是一个空链表,给 emp 找到对应的位置并插入
	// 让 cur 和 emp 比较，然后让 pre 保持在 cur 前面
	for {
		if cur != nil {
			// 比较
			if cur.Id > emp.Id {
				// 找到位置
				break
			}
			pre = cur // 保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	pre.Next = emp
	emp.Next = cur
}
//删除元素，找到链表位置，去掉节点
func (this *EmpLink) DelNode(id int) *Emp  {
	t := this.Head
	for {
		//val是头指针值得情况
		if t.Id == id {
			return t.Next
		}
		//循环结束条件
		if t.Next == nil {
			break
		}
		//val不是头指针值得情况
		if t.Next.Id == id {
			t.Next = t.Next.Next
			break
		}
		t = t.Next
	}
	return t
}

// 显示链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	// 遍历当前的链表，并显示数据
	cur := this.Head // 辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() // 换行处理
}

// 根据 id 查找对应的雇员，如果没有就返回 nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// hashtable ,含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 给 HashTable 编写 Insert 雇员的方法.
func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	// 使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}


// 显示 hashtable 的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

// 编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对于的链表的下标
}

// 完成查找
func (this *HashTable) FindById(id int) *Emp {
	// 使用散列函数，确定将该雇员应该在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}
// 完成查找
func (this *HashTable) DelById(id int)  {
	// 使用散列函数，确定将该雇员应该在哪个链表
	linkNo := this.HashFun(id)
	this.LinkArr[linkNo].Head = this.LinkArr[linkNo].DelNode(id)
}

func main() {
	var hashtable HashTable
	for i:=0;i<10 ;i++  {
		emp := &Emp{
			Id:   i,
			Name: "xxx"+strconv.Itoa(i),
			Next: nil,
		}
		hashtable.Insert(emp)
	}
	hashtable.ShowAll()
	hashtable.DelById(8)
	hashtable.ShowAll()
	emp1 := hashtable.FindById(2)
	emp1.ShowMe()

}