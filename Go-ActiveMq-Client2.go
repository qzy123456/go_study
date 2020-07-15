package main


import stomp "github.com/go-stomp/stomp"
import "fmt"

//Connect to ActiveMQ and listen for messages
func main() {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}

	sub, err := conn.Subscribe("/queue/test-1", stomp.AckMode(stomp.AckClientIndividual))
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg := <-sub.C
		fmt.Println(string(msg.Body),msg.ShouldAck())
        //如果需要应答
        if msg.ShouldAck() {
			conn.Ack(msg)
		}

	}
    //无消息超时
	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Disconnect()
}