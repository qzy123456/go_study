package main

import "../../rabbitmq20191121-master/RabbitMq"

func main() {
	two := RabbitMq.NewRabbitMqRouting("duExchangeName", "two")
	two.ConsumerRouting()
}
