package main

import "fmt"

func bar() (r int) {
	defer func() {
		r += 4
		if err :=recover();err != nil {
			r += 8
		}
	}()

	var f func() //// f没有初始化赋值，默认值是nil
	defer f() // 函数变量f的值已经确定下来是nil了,所以会触发panic，4+8+1=13
	//如果被defer的函数或方法的值是nil，在执行defer这条语句的时候不会报错，但是最后调用nil函数或方法的时候就引发panic:
	// runtime error: invalid memory address or nil pointer dereference。
	f = func() {
		fmt.Println("打印看看这里有没有执行到") //这里其实没有执行到
		r += 2
	}
	//defer f() //如果defer函数放在这里，函数正常执行，不会panic，所以 2+4+1 = 7
	return 1
}

func main() {
	println(bar()) //13
}