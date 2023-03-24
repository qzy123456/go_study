package main

import "rabbitmq20191121-master/RabbitMq"

func main() {
	one := RabbitMq.NewRabbitMqRouting("duExchangeName", "one")
	one.ConsumerRouting()
}
