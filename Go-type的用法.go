package main

import "fmt"

//使用type定义函数 func(int)
//定义 Call 接口，Call中有一个函数 call(int)
//在main()中调用one(2, callback)，在one()中调用two()，传入two()函数前，
//对callback函数实现了类型转换，从普通函数转换成type定义的函数。
//在 two() 中调用传入的 c 因为 c 实现了 Call 接口，所以可以调用 call() 函数，最终调用了我们传入的 callback() 函数。
func main() {
	one(2, callback)
}

//需要传递函数
func callback(i int) {
	fmt.Println("i am callBack")
	fmt.Println(i)
}

//main中调用的函数
func one(i int, f func(int)) {
	two(i, fun(f))
}

//one()中调用的函数
func two(i int, c Call) {
	c.call(i)
}

//定义的type函数
type fun func(int)

//fun实现的Call接口的call()函数
func (f fun) call(i int) {
	f(i)
}

//接口
type Call interface {
	call(int)
}
