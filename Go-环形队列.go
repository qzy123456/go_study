package main

import (
	"fmt"
	"errors"
	"os"
)

//管理环形队列的结构
type Queue struct {
	maxSize int
	array [5]int
	head int
	tail int
}

//入队列
func (this *Queue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满！")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队列
func (this *Queue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列为空！")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

//显示队列
func (this *Queue) Show() {
	if this.IsEmpty() {
		fmt.Println("队列为空！")
	}
	temp := this.head
	for i := 0; i < this.Size(); i++ {
		fmt.Printf("array[%d]:%d\t",temp,this.array[temp])
		temp = (temp + 1) % this.maxSize
	}

}

//判断队列是否已满
func (this *Queue) IsFull() bool {
	return (this.tail + 1) % this.maxSize == this.head
}

//判断队列是否为空
func (this *Queue) IsEmpty() bool {
	return this.head == this.tail
}

//查询有多少个队列
func (this *Queue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}




func main(){

	quque := &Queue {
		maxSize : 5,
		head : 0,
		tail : 0,
	}

	var xz string
	var number int
	for {
		fmt.Println()
		fmt.Println("1.添加队列请输入add")
		fmt.Println("2.获取队列请输入get")
		fmt.Println("3.显示队列请输入show")
		fmt.Println("4.输入exit退出")
		fmt.Scanln(&xz)

		switch xz {
		case "add" :
			fmt.Println("输入你要入列的数：")
			fmt.Scanln(&number)
			err := quque.Push(number)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("加入队列成功!\n")
			}
		case "get" :
			val, err := quque.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("队列已取出：%d",val)
			}

		case "show" :
			quque.Show()
			fmt.Println()
		case "exit" :
			os.Exit(0)
		}
	}
}