package main
import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	c, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Fprintf(os.Stdout, "sarama.NewConsumer err, message=%v \n", err)
		return
	}
	defer c.Close()
   //这里sarama.OffsetOldest是设置偏移量，也就是从哪条信息开始读取。默认是从0开始
   //可以根据具体情况（中间断线了）  记录最后一次的偏移量 下次从这个开始
	cp, err := c.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Fprintf(os.Stdout, "try create partition_consumer err, message=%v \n", err)
		return
	}
	defer cp.Close()

	for {
		select {
		case msg := <-cp.Messages():
			fmt.Fprintf(os.Stdout, "msg offset: %d, partition: %d, timestamp: %d, value: %s \n", msg.Offset, msg.Partition, msg.Timestamp.Unix(), string(msg.Value))
		case err := <-cp.Errors():
			fmt.Fprintf(os.Stdout, "err :%v \n", err)
		}
	}
}
