package main

import (
	"fmt"
)

type A struct {
	Text string
	Name string
}

func (a *A) Say() {
	fmt.Printf("A::Say():%s\n", a.Text)
}

type B struct {
	A
	Name string
}

func (b *B) Say() {
	b.A.Say()
	fmt.Printf("B::Say():%s\n", b.Text)
}

func main() {
	b := B{A{"hello, world", "张三"}, "李四"}
	//只要是方法里面是指针传递的话，那么是不是指针调用都是可以的
	//b := &B{A{"hello, world", "张三"}, "李四"}

	b.Say()
	fmt.Println("b的名字为：", b.Name)

	// 如果要显示 B 的 Name 值
	fmt.Println("b的名字为：", b.A.Name)
}