package main

import (
	"fmt"
)


func main() {
	array := []int{10, 11, 12, 13, 14}
	slice := array[0:4] // slice是对array的引用
	fmt.Println("array: ", array) // array:  [10 11 12 13 14]
	fmt.Println("slice: len",len(slice),"slice: cap=", cap(slice), ", value=", slice) // len 4 slice: cap= 5 , value= [10 11 12 13]


	array[0] += 10 // 会同时修改slice[0]
	slice[1] += 10 // 会同时修改array[1]
	fmt.Println("\nafter add 10")
	fmt.Println("array: ", array) // array:  [20 21 12 13 14]
	fmt.Println("slice: ", slice) // slice:  [20 21 12 13]


	slice1 := append(slice, 15) // 增加新元素15, cap仍然为5，array[4]变成15
	fmt.Println("\nafter append 15")
	fmt.Println("array: ", array) // array:  [20 21 12 13 15]
	fmt.Println("slice: ", slice) // slice:  [20 21 12 13]
	fmt.Println("slice: len",len(slice1),"slice1: cap=", cap(slice1), ", value=", slice1) // len 5, slice1: cap= 5 , value= [20 21 12 13 15]


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
	fmt.Println("slice2：cap=", cap(slice2), ", value=", slice2) // slice2：cap= 10 , value= [20 21 32 33 35 16]


	array[0] += 30 // 修改array[0]的值, slice[0]、slice1[0]的值会变化，但slice2[0]的值不变
	slice[1] += 30 // 修改slice[1]的值, array[1]、slice1[1]的值会变化，但slice2[1]的值不变
	slice1[2] += 30 // 修改slice1[2]的值, array[2]、slice[2]的值会变化，但slice2[2]的值不变
	slice2[3] += 30 // 修改slice2, array、slice、slice1的值未变化
	fmt.Println("\nafter add 30")
	fmt.Println("array: ", array) // array:  [50 51 62 33 35]
	fmt.Println("slice: ", slice) // slice:  [50 51 62 33]
	fmt.Println("slice1: ", slice1) // slice1:  [50 51 62 33 35]
	fmt.Println("slice2: ", slice2) // slice2:  [20 21 32 63 35 16]



}

