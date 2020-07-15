package main

import (
	"fmt"
	"time"
)
func main() {
	// 创建一个打点器, 每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)
	// 创建一个计时器, 2秒后触发
	stopper := time.NewTimer(time.Second * 2)
	// 声明计数变量
	var i int
	// 不断地检查通道情况
	for {
		// 多路复用通道
		select {
		case <-stopper.C:  // 计时器到时了
			fmt.Println("stop")
			// 跳出循环
			goto StopHere
		case <-ticker.C:  // 打点器触发了
			// 记录触发了多少次
			i++
			fmt.Println("tick", i)
		}
	}
	// 退出的标签, 使用goto跳转
StopHere:
	fmt.Println("done")
}
//tick 1
//tick 2
//tick 3
//tick 4
//stop
//done