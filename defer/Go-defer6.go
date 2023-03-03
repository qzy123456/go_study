package main

import "fmt"

func f() {
	defer func() {
		defer func() { fmt.Println("第二个",recover()) }()
		//fmt.Println("第二个",recover()) //没defer就是另一种
		panic(2)
	}()
	panic(1)
}

func main() {
	defer func() { fmt.Println("第一个",recover()) }()
	f()
	//第二个 2
	//第一个 1
}