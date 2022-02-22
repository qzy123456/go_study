package main

import (
	"fmt"
	"unsafe"
)

type W struct {
	b int32
	c int64
}
//unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
//而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
//unsafe.Pointer 可以和 普通指针 进行相互转换；
//unsafe.Pointer 可以和 uintptr 进行相互转换。
func main() {
	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b,w.c)

	//现在我们通过指针运算给b变量赋值为10
	b1 := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b1)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b,w.c)
	//uintptr(unsafe.Pointer(w)) 获取了 w 的指针起始值
	//unsafe.Offsetof(w.b) 获取 b 变量的偏移量
	//两个相加就得到了 b 的地址值，将通用指针 Pointer 转换成具体指针 ((*int)(b))，
	// 通过 * 符号取值，然后赋值。*((*int)(b)) 相当于把 (*int)(b) 转换成 int 了，
	// 最后对变量重新赋值成 10，这样指针运算就完成了
	var a *int8
	var b *int16
	a = new(int8)
	b = new(int16)
	*b = 10
	upb := unsafe.Pointer(b)
	fmt.Println(b,upb)
	b_int8ptr := (*int8)(upb)
	*a = *(b_int8ptr)
	fmt.Println(*a)
	//初始化一个指针类型的int8 a
	//初始化一个指针类型的int16 b
	//给指针b指向的内存赋值10
	//通过unsafe.pointer获取b的指针upb
	//把upb转成*int8指针 b_int8ptr
	//获取b_int8ptr的内存值赋值给a的地址指向的空间
}