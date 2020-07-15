package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
func worker(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <- taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}


func main() {

	rand.Seed(time.Now().UnixNano())

	const Max= 100
	const NumReceivers= 10
	const NumSenders= 10

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// It must be a buffered channel.
	//toStop 就是中间人的角色，使用它来接收 senders 和 receivers 发送过来的关闭 dataCh 请求。
	toStop := make(chan string, 1)

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				fmt.Println("this is random number of",value)
				if value == 0 {
					select {
						case toStop <- "sender#" + id:
						default:
					}
					return
				}

				select {
					case <-stopCh:
						return
					case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
							case toStop <- "receiver#" + id:
							default:
						}
						return
					}

					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}




	//////////////////
	taskCh := make(chan int, 100)
	go worker(taskCh)

	// 塞任务
	for i := 0; i < 10; i++ {
		taskCh <- i
	}

	// 等待 1 小时
	select {
		case <-time.After(time.Minute):
	}

}