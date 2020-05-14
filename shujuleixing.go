package main

import (
	"fmt"
	"math"
)
func main() {
	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", a, a)
	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)
	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))
	aa := []int{1, 2, 3}
	fmt.Println(aa[:2])
	//make()制造切片
	f:= make([]int, 2)
	d := make([]int, 20, 100)
	fmt.Println(f, d)
	fmt.Println(len(f), len(d))
    //append扩容
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
	//append() 函数除了添加一个元素外，也可以一次性添加很多元素。
	var car []string

	// 添加1个元素
	car = append(car, "OldDriver")

	// 添加多个元素
	car = append(car, "Ice", "Sniper", "Monk")
	// 添加切片
	team := []string{"Pig", "Flyingcake", "Chicken"}
	car = append(car, team...)
	fmt.Println(car)
}