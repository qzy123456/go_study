package main

import (
	"fmt"
	"time"
)

func main()  {

	ch := make(chan int)
	quit := make(chan bool)

	// 新开一个协程
	go func() {
		for ; ;  {
			select {
			case v := <-ch:
				fmt.Println(v)
			case <-time.After(3*time.Second):
				fmt.Println("Timeout.")
				quit<-true
				break
			}
		}
	}()

	//往ch中存放数据
	for i:=0;i<5;i++{
		ch<-i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("It is the end of the program.")
}