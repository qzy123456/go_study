package main

import "fmt"

func test1() int {
	var i = 0
	defer func() {
		i = 10
	}()
	return i
}

func test2() (result int) {
	defer func() {
		result *= 10
	}()
	return 6
}

func test3() (result int) {
	result = 8
	defer func() {
		result *= 10
	}()
	return
}

func main() {
	result1 := test1() //0
	result2 := test2() //60
	result3 := test3() //80
	fmt.Println(result1, result2, result3)
}