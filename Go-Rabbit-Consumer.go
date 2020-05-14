package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError2(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 只能在安装 rabbitmq 的服务器上操作
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError2(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError2(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError2(err, "Failed to declare a queue")
//公平调度（Fair dispatch）
	//RabbitMQ的默认消息分配不能够满足我们的需要，比如有两个消费者，其中一个消费者经常忙碌的状态，
	// 另外一个消费者几乎不做任何工作，但是RabbitMQ仍然均匀的在两者之间调度消息。
	// 这是因为RabbitMQ只做队列当中的消息调度而没有查看某个消费者中未确认的消息，它只是盲目的将第n条消息发送给第n个消费者
	//解决以上问题我们可以设置prefetch count数值为1，这样只有当消费者消费完消息并返回ack确认后RabbitMQ才会给其分发消息，
	// 否则只会将消息分发给其他空闲状态的消费者
	//注意：消费者需要的话，要设置，生产者不用设置
	err = ch.Qos(
		1, //// 在没有返回ack之前，最多只接收1个消息
		0,
		false,
	)

//queue:队列名称。
	//consumer:消费者标签，用于区分不同的消费者。
	//autoAck:是否自动回复ACK，true为是，回复ACK表示高速服务器我收到消息了。建议为false，手动回复，这样可控性强。
	//exclusive:设置是否排他，排他表示当前队列只能给一个消费者使用。
	//noLocal:如果为true，表示生产者和消费者不能是同一个connect。
	//nowait：是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ Server的返回信息，而RMQ Server也不会返回信息。（不推荐使用）
	//args：直接写nil，没研究过，不解释。
	//注意下返回值：返回一个<- chan Delivery类型，遍历返回值，有消息则往下走， 没有则阻塞。
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // //将autoAck设置为false，则需要在消费者每次消费完成
							// 消息的时候调用d.Ack(false)来告诉RabbitMQ该消息已经消费
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError2(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			//multiple为true的时候：此次交付和之前没有确认的交付都会在通过同一个通道交付，这在批量处理的时候很有用
			//为false的时候只交付本次。只有该方法执行了，RabbitMQ收到该确认才会将消息删除
			d.Ack(false)
		}

	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
