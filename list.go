package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	//fmt.Println(l)
	//使用 for 语句进行遍历，其中 i:=l.Front() 表示初始赋值，只会在一开始执行一次；每次循环会进行一次 i!=nil 语句判断，
	// 如果返回 false，表示退出循环，反之则会执行 i=i.Next()。
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	// 使用
	//l.Remove(element)
}