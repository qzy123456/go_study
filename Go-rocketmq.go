package main

import (
	"github.com/zhucheer/orange/queue"
	"time"
	"fmt"
)

func main(){
	// 注册生产者 填入broker节点,group名称,重试次数信息
	mqProducerClient := queue.RegisterRocketProducerMust([]string{"192.168.137.100:9876"}, "test", 1)

	// 注册消费者 填入broker节点,group名称信息
	mqConsumerClient:= queue.RegisterRocketConsumerMust([]string{"192.168.137.100:9876"}, "test")

	go func() {
		for i:=0;i<10;i++{
			// 向队列发送一条消息 填入消息队列topic和消息体信息
			ret,_:=mqProducerClient.SendMsg("topicTest", "Hello mq~~")
			fmt.Println("========producer push one message====", ret.MsgId)

			time.Sleep(time.Second)
		}

	}()

	// 执行消费者监听 填入消息队列topic
	mqConsumerClient.ListenReceiveMsgDo("topicTest", func(mqMsg queue.MqMsg) {
		// 收到一条消息
		fmt.Println("receive====>",mqMsg.MsgId, mqMsg.BodyString())

	})

	time.Sleep(20*time.Second)
}

