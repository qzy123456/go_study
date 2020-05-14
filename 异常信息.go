package main

import "fmt"

func main() {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err.(string)) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()
	f(1, 0) //开始调用f
	data := [...]int{0, 1, 2, 3, 4, 5, 6}
	slice := data[1:4:5] //// [low : high : max]
	fmt.Println(slice)
	panic("212")
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f(x, y int) {
	var z int
	func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err) // 这里的err其实就是panic传入的内容
			}
		}()
		z = x / y
		return
	}()
	println("x / y =", z)

}