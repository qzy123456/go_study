package main

import (
	"fmt"
	"reflect"
)
var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)
func Make(T reflect.Type, fptr interface{}) {
	// 实际创建 slice 的包装函数。
	swap := func(in []reflect.Value) []reflect.Value {

		// 返回和类型匹配的 slice 对象。
		return []reflect.Value{
			reflect.MakeSlice(
				reflect.SliceOf(T),// slice type
				int(in[0].Int()),// len
				int(in[1].Int()),// cap
			),
		}
	 }
	fn := reflect.ValueOf(fptr).Elem()// 传⼊的是函数变量指针，因为我们要将变量指向 swap 函数。
	v := reflect.MakeFunc(fn.Type(), swap)// 获取函数指针类型，⽣生成所需 swap function value。
	fn.Set(v) // 修改函数指针实际指向，也就是 swap。
}

func main() {
	var makeints func(int, int) []int
	var makestrings func(int, int) []string
	// ⽤相同算法，⽣成不同类型创建函数。
	 Make(Int, &makeints)
	 Make(String, &makestrings)
	// 按实际类型使⽤用。
	x := makeints(5, 10)
	fmt.Printf("%#v\n", x)
	s := makestrings(3, 10)
	fmt.Printf("%#v\n", s)
	//输出 打印
	//[]int{0, 0, 0, 0, 0}
	//[]string{"", "", ""}
}