package main

import "rabbitmq20191121-master/RabbitMq"

func main() {
	all := RabbitMq.NewRabbitMqTopic("exchangeNameTpoic1224", "#")
	all.ConsumerTopic()
}
