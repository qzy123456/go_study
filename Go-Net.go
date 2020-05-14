package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string){
	for _,data:=range str {
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

//person1执行完成，才到person2执行
func person1(){
	Printer("Oh!")
	ch<-0    //给管道/通道写数据，发送
}

func person2(){
	<-ch    //从管道取数据，接收，如果通道没有数据它就会阻塞
	Printer("Yeah!")
}

func main() {

	//新建2个协程，代表2个人。两个人共同使用打印机
	go person1()
	go person2()

	//不让主协程结束
	for{}
}