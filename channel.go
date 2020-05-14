package main

import (
        "fmt"
)

func sum(s []int, c chan int) {
        sum := 0
        for _, v := range s {
                sum += v
        }
        c <- sum // 把 sum 发送到通道 c
}
func onlyW(w chan<- int) {
        defer close(w)
        for i:=0;i<10;i++ {
                // 只能向w中写入数据
                w<-i
        }
}
func onlyR(r <-chan int){
        for data := range  r{
                fmt.Printf("%d\t",data)
        }
}
func main() {
        s := []int{7, 2, 8, -9, 4, 0,1,2}

        c := make(chan int)
        go sum(s[:len(s)/2], c)
        go sum(s[len(s)/2:], c)
        x, y := <-c, <-c // 从通道 c 中接收

        fmt.Println(x, y, x+y)
        // 定义双向channel变量ch1
        ch1 := make(chan int, 10)
        // 隐式转换成单向 只写channel
        var send chan<- int = ch1
        // 隐式转换成单向 只读channel
        var rece <-chan int = ch1
        go onlyW(send)
        onlyR(rece)
        fmt.Println("it is over")
}
