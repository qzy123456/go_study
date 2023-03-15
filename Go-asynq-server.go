package main

import (
	"fmt"
	"encoding/json"
    "context"
	"github.com/hibiken/asynq"
)

var AsynqServer *asynq.Server // kodo中的异步任务server

func main() {
	err := initTaskServer()
	if err != nil{
		fmt.Println(err)
	}
	mux := asynq.NewServeMux()

	mux.HandleFunc("test99", HandleUpdateTask)
	// ...register other handlers...

	if err := AsynqServer.Run(mux); err != nil {
		fmt.Printf("could not run asynq server: %v", err)
	}
}

func HandleUpdateTask(ctx context.Context, t *asynq.Task) error {
	res := make(map[string]string)

	err := json.Unmarshal(t.Payload(), &res)
	if err != nil {
		fmt.Printf("rum session, can not parse payload: %s,  err: %v", t.Payload(), err)
		return nil
	}
	//--------具体处理逻辑------------//
	fmt.Printf("setting:%+#v\n", res)
	return nil
}


func initTaskServer() error {
	// 初始化异步任务服务端
	AsynqServer = asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     "127.0.0.1:6379",
			Password: "", //与client对应
			DB:       0,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 100,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)
	return nil
}
