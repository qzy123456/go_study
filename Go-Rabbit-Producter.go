package main

import (
"log"

"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 只能在安装 rabbitmq 的服务器上操作
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
//name:交换器的名称，对应图中exchangeName。
	//kind:也叫作type，表示交换器的类型。有四种常用类型：direct、fanout、topic、headers。
	//durable:是否持久化，true表示是。持久化表示会把交换器的配置存盘，当RMQ Server重启后，会自动加载交换器。
	//autoDelete:是否自动删除，true表示是。至少有一条绑定才可以触发自动删除，当所有绑定都与交换器解绑后，会自动删除此交换器。
	//exclusive：是否设置排他，true为是。如果设置为排他，则队列仅对首次声明他的连接可见，并在连接断开时自动删除。（注意，这里说的是连接不是信道，相同连接不同信道是可见的）。
	//internal:是否为内部，true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
	//noWait:是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ Server的返回信息，而RMQ Server也不会返回信息。（不推荐使用）
	//args:直接写nil，没研究过，不解释。
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	//exchange：要发送到的交换机名称，
	//key：路由键，对应图中RoutingKey。
	//mandatory：直接false，不建议使用，后面有专门章节讲解。
	//immediate ：直接false，不建议使用，后面有专门章节讲解。
	//msg：要发送的消息，msg对应一个Publishing结构，Publishing结构里面有很多参数，这里只强调几个参数，其他参数暂时列出，但不解释。
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}