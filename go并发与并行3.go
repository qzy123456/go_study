package main

import (
	"fmt"
	"runtime"
)
var quit2 chan int = make(chan int)

func loop2() {
	for i := 0; i < 100; i++ { //为了观察，跑多些
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		fmt.Printf("%d ", i)
	}
	quit2 <- 0
}

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(4) // 最多使用2个核

	go loop2()
	go loop2()

	for i := 0; i < 2; i++ {
		<- quit2
	}
}