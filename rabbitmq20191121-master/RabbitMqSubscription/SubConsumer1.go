package main

import "../../rabbitmq20191121-master/RabbitMq"

func main() {
	rabbitmq := RabbitMq.NewRabbitMqSubscription("duexchangeName")
	rabbitmq.ConsumeSbuscription()
}
