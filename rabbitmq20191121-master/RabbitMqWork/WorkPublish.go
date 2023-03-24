package main

import (
	"fmt"
	"rabbitmq20191121-master/RabbitMq"
	"strconv"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("duQueueName191224")
	for i := 0; i < 10; i++ {
		rabbitmq.PublishSimple("hello du message" + strconv.Itoa(i) + "---来自work模式")
		fmt.Printf("work模式，共产生了%d条消息\n", i)
	}
}
