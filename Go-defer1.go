package main

func bar() (r int) {
	defer func() {
		r += 4
		if err :=recover();err != nil {
			r += 8
		}
	}()

	var f func() //// f没有初始化赋值，默认值是nil
	defer f() // 函数变量f的值已经确定下来是nil了,所以会触发panic
	//如果被defer的函数或方法的值是nil，在执行defer这条语句的时候不会报错，但是最后调用nil函数或方法的时候就引发panic:
	// runtime error: invalid memory address or nil pointer dereference。
	f = func() {

		r += 2
	}

	return 1
}

func main() {
	println(bar())
}