package main


import (
	"fmt"
	"github.com/go-stomp/stomp"
	"time"
)

func recv_data(ch chan *stomp.Message) {
	//不断地循环，从channel里面获取数据
	for {
		v := <-ch
		//这里是打印当然还可以做其他的操作，比如写入hdfs平台
		//v是*stomp.Message类型，属性都在这里面
		fmt.Println(string(v.Body))
	}
}

func main() {
	//创建一个channel，存放的是*stomp.Message类型
	ch := make(chan *stomp.Message)
	//将管道传入函数中
	go recv_data(ch)
	//和生产者一样，调用Dial方法，返回conn和err
	conn, err := stomp.Dial("tcp", "127.0.0.1:61613")
	if err != nil {
		fmt.Println("err =", err)
	}
	//消费者订阅这个队列
	//参数一：队列名
	//参数二：确认信息，直接填默认地即可
	sub, err := conn.Subscribe("testQ", stomp.AckMode(stomp.AckAuto))
	for { //无限循环
		select {
		//sub.C是一个channel，如果订阅的队列有数据就读取
		case v := <-sub.C:
			//读取的数据是一个*stomp.Message类型
			ch <- v
			//如果30秒还没有人发数据的话，就结束
		case <-time.After(time.Second * 30):
			return
		}
	}
}
