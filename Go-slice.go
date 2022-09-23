package main

import (
	"fmt"
)
func main() {
	array := []int{10, 11, 12, 13, 14}
	slice := array[0:4] // slice是对array的引用
	fmt.Println("array: ", array) // array:  [10 11 12 13 14]
	fmt.Println("slice: len",len(slice),"slice: cap=", cap(slice), ", value=", slice) 
	//slice: len 4 slice: cap= 5 , value= [10 11 12 13]

	array[0] += 10 // 会同时修改slice[0]
	slice[1] += 10 // 会同时修改array[1]
	fmt.Println("\nafter add 10")
	fmt.Println("array: ", array) // array:  [20 21 12 13 14]
	fmt.Println("slice: ", slice) // slice:  [20 21 12 13]

	slice1 := append(slice, 15) // 增加新元素15, cap仍然为5，array[4]变成15
	fmt.Println("\nafter append 15")
	fmt.Println("array: ", array) // array:  [20 21 12 13 15]
	fmt.Println("slice: ", slice) // slice:  [20 21 12 13]
	fmt.Println("slice: len",len(slice1),"slice1: cap=", cap(slice1), ", value=", slice1) 
	// len 5, slice1: cap= 5 , value= [20 21 12 13 15]

	array[2] += 20 // 会同时修改slice[2]、slice1[2]
	slice[3] += 20 // 会同时修改array[3]、slice1[3]
	slice1[4] += 20 // 会同时修改array[4]
	fmt.Println("\nafter add 20")
	fmt.Println("array: ", array) // array:  [20 21 32 33 35]
	fmt.Println("slice: ", slice) // slice:  [20 21 32 33]
	fmt.Println("slice1: ", slice1) // slice1: [20 21 32 33 35]


	slice2 := append(slice1, 16) // 添加新元素16，cap变为10，array的值未变化
	fmt.Println("\nafter append 16")
	fmt.Println("array: ", array) // array:  [20 21 32 33 35]
	fmt.Println("slice: ", slice) // slice:  [20 21 32 33]
	fmt.Println("slice1: ", slice1) // slice1: [20 21 32 33 35]
	fmt.Println("slice2：cap=", cap(slice2), ", value=", slice2) 
	// slice2：cap= 10 , value= [20 21 32 33 35 16]

	array[0] += 30 // 修改array[0]的值, slice[0]、slice1[0]的值会变化，但slice2[0]的值不变
	slice[1] += 30 // 修改slice[1]的值, array[1]、slice1[1]的值会变化，但slice2[1]的值不变
	slice1[2] += 30 // 修改slice1[2]的值, array[2]、slice[2]的值会变化，但slice2[2]的值不变
	slice2[3] += 30 // 修改slice2, array、slice、slice1的值未变化
	fmt.Println("\nafter add 30")
	fmt.Println("array: ", array) // array:  [50 51 62 33 35]
	fmt.Println("slice: ", slice) // slice:  [50 51 62 33]
	fmt.Println("slice1: ", slice1) // slice1:  [50 51 62 33 35]
	fmt.Println("slice2: ", slice2) // slice2:  [20 21 32 63 35 16]
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	fmt.Println(len(s),cap(s))
	x := append(s, 11)
	fmt.Println(len(x),cap(x))
	y := append(s, 12)
	fmt.Println(len(y),cap(y))
	fmt.Println(s, x, y)
	//[5 7 9] [5 7 9 12] [5 7 9 12]
	//s := []int{5}       s 只有一个元素，[5]
	//s = append(s, 7)    s 扩容，容量变为2，[5, 7]
	//s = append(s, 9)    s 扩容，容量变为4，[5, 7, 9]。注意，这时 s 长度是3，只有3个元素
	//x := append(s, 11)  由于 s 的底层数组仍然有空间，因此并不会扩容。这样，底层数组就变成了 [5, 7, 9, 11]。
						  //注意，此时 s = [5, 7, 9]，容量为4；x = [5, 7, 9, 11]，容量为4。这里 s 不变
	//y := append(s, 12)  这里还是在 s 元素的尾部追加元素，由于 s 的长度为3，容量为4，所以直接在底层数组索引为3的地方填上12。
	   				      //结果：s = [5, 7, 9]，y = [5, 7, 9, 12]，x = [5, 7, 9, 12]，x，y 的长度均为4，容量也均为4
	//所以最后程序的执行结果是：
	//[5 7 9] [5 7 9 12] [5 7 9 12]
	//这里要注意的是，append函数执行完后，返回的是一个全新的 slice，并且对传入的 slice 并不影响。
}

