package main

import "fmt"
//变参本质上就是 slice。只能有⼀一个，且必须是最后⼀一个。
func test(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
//变参本质上就是 slice。只能有⼀一个，且必须是最后⼀一个。
func main()  {
	//使⽤用 slice 对象做变参时，必须展开
	println(test("sum: %d", 1, 2, 3))
	s := []int{1, 2, 3}
	println(test("sum: %d", s...))
	//命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
	println(add(1, 2))
	//命名返回参数允许 defer 延迟调⽤用通过闭包读取和修改。
    println(add1(1,3))
}
//命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
func add(x, y int) (z int) {
	z=x+y
	return
	}
//命名返回参数允许 defer 延迟调⽤用通过闭包读取和修改。
func add1(x, y int) (z int) {
	defer func() {
		z += 100
		}()
	z=x+y
	return
	}