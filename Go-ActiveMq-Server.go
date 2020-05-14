package main
import (
	"fmt"
	"github.com/go-stomp/stomp"
	"time"
)

func main(){
	// 调用Dial方法，第一个参数是"tcp"，第二个参数则是ip:port
	// 返回conn(连接)和err(错误)
	conn,err:=stomp.Dial("tcp", "localhost:61613")
	// 错误判断
	if err!=nil{
		fmt.Println("err =", err)
		return
	}
	//发送十条数据
	for i:=0;i<10;i++ {
		// 调用conn下的send方法，接收三个参数
		//参数一：队列的名字
		//参数二：数据类型，一般是文本类型，直接写text/plain即可
		//参数三：内容，记住要转化成byte数组的格式
		//返回一个error
		err := conn.Send("testQ", "text/plain",[]byte(fmt.Sprintf("message:%d", i)))
		if err!=nil{
			fmt.Println("err =", err)
		}
	}
	/*
	这里为什么要sleep一下，那就是conn.Send这个过程是不阻塞的
	相当于Send把数据放到了一个channel里面
	另一个goroutine从channel里面去取数据再放到消息队列里面
	但是还没等到另一个goroutine放入数据，此时循环已经结束了
	因此最好要sleep一下，根据测试，如果不sleep，那么发送1000条数据，
	最终进入队列的大概是980条数据，这说明了什么
	说明了当程序把1000条数据放到channel里面的时候，另一个goroutine只往队列里面放了980条
	剩余的20条还没有来得及放入，程序就结束了
	 */
	time.Sleep(time.Second * 1)
}
