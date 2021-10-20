package main

import "fmt"

func main()  {

	// 指针类型的nil比较
	fmt.Println((*int64)(nil) == (*int64)(nil))
	// channel 类型的nil比较
	fmt.Println((chan int)(nil) == (chan int)(nil))
	// func类型的nil比较
	fmt.Println((func())(nil) == (func())(nil)) // func() 只能与nil进行比较
	// interface类型的nil比较
	fmt.Println((interface{})(nil) == (interface{})(nil))
	// map类型的nil比较
	fmt.Println((map[string]int)(nil) == (map[string]int)(nil)) // map 只能与nil进行比较
	// slice类型的nil比较
	fmt.Println(([]int)(nil) == ([]int)(nil)) // slice 只能与nil进行比较
    //从运行结果我们可以看出，指针类型nil、channel类型的nil、interface类型可以相互比较，
    //而func类型、map类型、slice类型只能与nil标识符比较，两个类型相互比较是不合法的。
	var ptr *int64 = nil
	var cha chan int64 = nil
	var fun func() = nil
	var inter interface{} = nil
	var ma map[string]string = nil
	var slice []int64 = nil
	fmt.Println(ptr == cha)
	fmt.Println(ptr == fun)
	fmt.Println(ptr == inter)
	fmt.Println(ptr == ma)
	fmt.Println(ptr == slice)

	fmt.Println(cha == fun)
	fmt.Println(cha == inter)
	fmt.Println(cha == ma)
	fmt.Println(cha == slice)

	fmt.Println(fun == inter)
	fmt.Println(fun == ma)
	fmt.Println(fun == slice)

	fmt.Println(inter == ma)
	fmt.Println(inter == slice)

	fmt.Println(ma == slice)
	//从运行结果我们可以得出，只有指针类型和channel类型与接口类型可以比较，其他类型的之间是不可以相互比较的。
	//为什么指针类型、channel类型可以和接口类型进行比较呢？
}