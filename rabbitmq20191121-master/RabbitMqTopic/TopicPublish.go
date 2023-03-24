package main

import (
	"fmt"
	"rabbitmq20191121-master/RabbitMq"
	"strconv"
)

func main() {
	one := RabbitMq.NewRabbitMqTopic("exchangeNameTpoic1224", "Singer.Jay")
	two := RabbitMq.NewRabbitMqTopic("exchangeNameTpoic1224", "Persident.XIDADA")
	for i := 0; i < 5; i++ {
		one.PublishTopic("小杜同学，topic模式，Jay," + strconv.Itoa(i))
		two.PublishTopic("小杜同学，topic模式，All," + strconv.Itoa(i))
		fmt.Printf("topic模式。这是小杜同学发布的消息%v \n", i)
	}
}
