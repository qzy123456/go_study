package main

import (
	"fmt"
	"rabbitmq20191121-master/RabbitMq"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("duQueueName1912161843")
	rabbitmq.ConsumeSimple()
	fmt.Println("接收成功！")
}
