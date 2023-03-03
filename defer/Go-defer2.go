package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("111")
		fmt.Println(recover()) //这个defer其实捕获的是下面panic(1)
	}()
	defer func() {
		fmt.Println("222")
		defer fmt.Println(recover()) //这个defer捕获的是panic(2)
		defer panic(1)
		fmt.Println("333")
		recover()
	}()
	defer recover()
	fmt.Println("444")
	panic(2)
	//444
	//222
	//333
	//2
	//111
	//1
}