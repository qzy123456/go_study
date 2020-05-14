package main

import (
	"fmt"
	"os"
)

func tests() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
	}
func test11() error {
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	defer f.Close() // 注册调⽤用，⽽而不是注册函数。必须提供参数，哪怕为空。
	f.WriteString("Hello, World!")
	return nil
}
//匿名函数可赋值给变量，做为结构字段，或者在 channel ⾥里传送。
func main()  {
	fn := func() {
		println("Hello, World!")
	}
	fn()

	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](111))
	println(fns[1](111))
	//println(fns[2](111))  超出范围
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	println(d.fn())
	//通道
	fc := make(chan func() string, 2)
	fc <- func() string {
		return "Hello, World！！!"
	}
	println((<-fc)())
	f := tests()
	f()
	//注册跳用，而不是注册函数，需要提供参数
	_ = test11()
}