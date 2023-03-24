package main

import (
	"fmt"
	"rabbitmq20191121-master/RabbitMq"
	"strconv"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMqSubscription("duexchangeName")
	for i := 0; i < 5; i++ {
		rabbitmq.PublishSubscription("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		fmt.Printf("订阅模式生产第" + strconv.Itoa(i) + "条数据\n")
	}
}
