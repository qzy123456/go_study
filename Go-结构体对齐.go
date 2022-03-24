package main

import (
	"fmt"
	"unsafe"
)

//对于结构体的每个字段，我们先来认识一下如下4个概念：
//对齐宽度
//本身所占字节数
//实际占用字节数
//偏移量
//那么他们之间有什么关系呢？
//对齐宽度≤本身所占字节数≤实际占用字节数。
//首先来说，本身所占字节数就是类型大小。当把类型放到结构体时，它实际占用字节数是大于等于类型本身大小的，
// 多出来的部分叫填充字节。也就是说实际占用字节数=本身所占字节数+填充字节数。
//对齐宽度是类型的一种属性，他和类型本身以及操作系统有关。
// 一般情况下，对齐宽度和类型大小是一致的。比如byte和bool类型的对齐宽度是1字节，int32类型对齐宽度是4字节。
// 那为什么对齐宽度又会小于类型大小呢？那是因为对齐宽度有一个上限，在32位系统上，对齐宽度最大为4字节，
// 因此，即便是int64类型，对齐宽度也是4字节，而不是8字节；相应的，在64位系统上，对齐宽度为8字节，
// 即使是string(本身占16字节)，对齐宽度也只有8字节。因此，以下的例子都是在64位系统上的结果。

type st struct {
	f1 int64
	f2 byte
}
type s2 struct {
	f1 byte
	f2 int64
}

func main() {
	var s st
	fmt.Println("st")
	fmt.Println("对齐宽度：", unsafe.Alignof(s)) //8
	fmt.Println("本身大小：", unsafe.Sizeof(s))  //16
	fmt.Println("f1")
	fmt.Println("对齐宽度：", unsafe.Alignof(s.f1))   //8
	fmt.Println("本身大小：", unsafe.Sizeof(s.f1))    //8
	fmt.Println("偏 移 量：", unsafe.Offsetof(s.f1)) //0
	fmt.Println("f2")
	fmt.Println("对齐宽度：", unsafe.Alignof(s.f2))   //1
	fmt.Println("本身大小：", unsafe.Sizeof(s.f2))    //1
	fmt.Println("偏 移 量：", unsafe.Offsetof(s.f2)) //8
	var s2 s2
	fmt.Println("s2")
	fmt.Println("对齐宽度：", unsafe.Alignof(s2)) //8
	fmt.Println("本身大小：", unsafe.Sizeof(s2))  //16
	fmt.Println("s2.f1")
	fmt.Println("对齐宽度：", unsafe.Alignof(s2.f1))   //1
	fmt.Println("本身大小：", unsafe.Sizeof(s2.f1))    //1
	fmt.Println("偏 移 量：", unsafe.Offsetof(s2.f1)) //0
	fmt.Println("s2.f2")
	fmt.Println("对齐宽度：", unsafe.Alignof(s2.f2))   //8
	fmt.Println("本身大小：", unsafe.Sizeof(s2.f2))    //8
	fmt.Println("偏 移 量：", unsafe.Offsetof(s2.f2)) //8
}

/*
st
对齐宽度： 8
本身大小： 16
f1
对齐宽度： 8
本身大小： 8
偏 移 量： 0
f2
对齐宽度： 1
本身大小： 1
偏 移 量： 8
s2
对齐宽度： 8
本身大小： 16
s2.f1
对齐宽度： 1
本身大小： 1
偏 移 量： 0
s2.f2
对齐宽度： 8
本身大小： 8
偏 移 量： 8
*/
