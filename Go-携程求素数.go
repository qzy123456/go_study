package main

import (
	"fmt"
	"time"
)
// 计算一个数字内包含了多少个素数
var num int
// 开启的协程数量
var routineNum int
// 向通道内写入数据
func DataW(ch1 chan int) {
	// 关闭通道
	defer close(ch1)
	for i := 0; i < num; i++ {
		ch1 <- i
	}
}
// 判断字数内的素数,写入通道内
func checkNnum(input chan int, store chan int, exit chan bool) {
	var check bool
	for {
		// 从通道内读取数据
		data, ok := <-input
		// 没有数据的时候结束循环
		if !ok {
			break
		}
		// 素数判断逻辑
		check = true
		for i := 2; i < data; i++ {
			// 能整除除了1和本身以外的数字,那就不是素数
			if data%i == 0 {
				check = false
				break
			}
		}
		// 将符合的素数写入通道中
		if check {
			store <- data
		}
	}
	// 向终止通道写入数据
	exit <- true
}
func closeStore(exit chan bool, store chan int) {
	for i := 0; i < routineNum; i++ {
		<-exit
	}
	close(store)
}
func main() {
	// 获取客户端输入
	fmt.Println("请输入需要被计算的数(测试其中的素数) :")
	_, _ = fmt.Scanln(&num)
	fmt.Println("请输入开启的协程数")
	_, _ = fmt.Scanln(&routineNum)
	start := time.Now()
	var Res []int
	inputCh := make(chan int, 2000)
	storeCh := make(chan int, 500)
	exitCh := make(chan bool, routineNum)
	go DataW(inputCh)
	for i := 0; i < routineNum; i++ {
		go checkNnum(inputCh, storeCh, exitCh)
	}
	go closeStore(exitCh, storeCh)
	for {
		data, ok := <-storeCh
		if !ok {
			break
		}
		Res = append(Res, data)
	}
	for k, v := range Res {
		fmt.Printf("%d\t", v)
		if k%20 == 0 && k != 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("\nspend : ", time.Now().Sub(start))

}
