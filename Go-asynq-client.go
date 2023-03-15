package main

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)
var client *asynq.Client
func main() {
	client = asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB: 0,
		})
	updateTask()              // 触发
	time.Sleep(2*time.Second)
}


func updateTask(){
	data := map[string]string{
		"aa": "11",
		"bb": "22",
	}

	payload, _ := json.Marshal(data)

	t1 := asynq.NewTask("test99", payload)

	info, err := client.Enqueue(t1, asynq.TaskID("sdf11111"))//添加taskID,如果任务没有被消费，则入队列失败
	if err != nil {
		fmt.Printf("err:%v,info:%v\n", err, info)
	}

	fmt.Printf(" [*] Successfully enqueued task: %+v", info)
}

