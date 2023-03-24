package main

import "rabbitmq20191121-master/RabbitMq"

func main() {
	jay := RabbitMq.NewRabbitMqTopic("exchangeNameTpoic1224", "Singer.*")
	jay.ConsumerTopic()
}
