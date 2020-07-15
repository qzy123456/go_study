package main

import (
	"fmt"
	_"go-common 2/library/syscall"
	_"os"
	_"os/signal"
	"time"
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
		//close(done) //关闭也是可以执行完的
		done <- 1
	}()
	<-done
   //生产者  消费者
   ch := make(chan int,64)

   go Producter(3,ch)
   go Producter(5,ch)

   go Consumer(ch)
   //睡眠退出
   time.Sleep(1*time.Second)
   //ctrl+c退出
   // sig := make(chan os.Signal,1)
    //signal.Notify(sig,syscall.SIGINT,syscall.SIGTERM)
    //fmt.Println("tuichu",<-sig)
}