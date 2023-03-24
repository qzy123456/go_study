package main

import (
	"fmt"
	"rabbitmq20191121-master/RabbitMq"
	"strconv"
)

func main() {
	rabbitmq1 := RabbitMq.NewRabbitMqRouting("duExchangeName", "one")
	rabbitmq2 := RabbitMq.NewRabbitMqRouting("duExchangeName", "two")
	rabbitmq3 := RabbitMq.NewRabbitMqRouting("duExchangeName", "three")
	for i := 0; i < 10; i++ {
		rabbitmq1.PublishRouting("路由模式one" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("路由模式two" + strconv.Itoa(i))
		rabbitmq3.PublishRouting("路由模式three" + strconv.Itoa(i))
		fmt.Printf("在路由模式下，routingKey为one,为two,为three的都分别生产了%d条消息\n", i)
	}
}
