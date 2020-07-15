package main

import (
	"fmt"
	"time"
)

var quit1 chan int

func foo(id int) {
	fmt.Println(id)
	time.Sleep(time.Second) // 停顿一秒
	quit1 <- 0 // 发消息：我执行完啦！
}
func main() {
	count := 1000
	quit1 = make(chan int, count) // 缓冲1000个数据

	for i := 0; i < count; i++ { //开1000个goroutine
		go foo(i)
	}

	for i :=0 ; i < count; i++ { // 等待所有完成消息发送完毕。
		<- quit1
	}
}