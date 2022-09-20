package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())  //第一个defer捕获panic（2）
	}()
	defer func() {
		defer func() { //这里不加defer 就是捕获panic（2），结果就是2 1
			defer fmt.Print(recover()) //捕获panic（1）
		}()
		panic(1)
	}()
	defer recover() //无效
	panic(2) // 1 2
	//首先，第一个defer执行，将recover()压入栈中，
	// 其次第二个defer执行压入栈的第二层，然后第二个defer中又嵌套一个defer，将嵌套的defer压入第二个defer的defer栈中第一层，
	// 继续往下执行，遇到panic，抛出panic(1)。遇到panic程序不会立刻上传，
	// 而是将defer执行完。根据defer知识。defer是最后再执行，第二个defer执行完继续往下执行，遇到panic(2)，抛出panic(2)，
	// 然后回去调用第9行的defer，然后调用panic(1)，然后panic(1)会调用第11行的recover被捕获并打印出1，
	// 最后继续执行第7行，panic(2)被捕获并打印2。最后输出1，2
}
