package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//生产者  消费者
//生产者  生成 factor整数倍的序列

func Producter(factor int,out chan <- int)  {

	for i:=0;;i++ {
		out <- i * factor
	}
}
//消费者
func Consumer(in  <-chan int)  {
	for v :=range in{
		fmt.Println(v)
	}
}

func main()  {
	done := make(chan int)
	go func() {
		println("111")
		//close(done) //关闭就会deadlock
		done <- 1
	}()
	<-done
   //生产者  消费者
   ch := make(chan int,64)

   go Producter(3,ch)
   go Producter(5,ch)

   go Consumer(ch)

   //ctrl+c退出
    sig := make(chan os.Signal,1)
    signal.Notify(sig,syscall.SIGINT,syscall.SIGTERM)
    fmt.Println("tuichu",<-sig)
}