package main

import (
	"fmt"
	"os"
)

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
func main() {
	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})

	var n int = 1
	fmt.Println(Rabit(n))
	fmt.Println(Fib(n))

	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	Pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", Pid)
}

func Rabit(n int) int {
	for {
		if n == 2 || n == 1 {
			return 1
		}
		return Rabit(n-1) + Rabit(n-2)
	}
}

//元组赋值
func Fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
