/*
unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
unsafe.Pointer 可以和 普通指针 进行相互转换；
unsafe.Pointer 可以和 uintptr 进行相互转换。
*/
package main

import (
	"fmt"
	"unsafe"
)

type W struct {
	b int32
	c int64
}

type iface struct {
	itab, data uintptr
}

func main() {
	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	d := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(d)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)

	var a interface{} = nil

	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)

	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}
