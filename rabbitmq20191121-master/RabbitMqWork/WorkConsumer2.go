package main

import "rabbitmq20191121-master/RabbitMq"

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("duQueueName191224")
	rabbitmq.ConsumeSimple()
}
