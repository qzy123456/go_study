package main

import (
	"fmt"
	"time"
	"context"
)
type ctxKey string
func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// monitor
	go func() {
		for range time.Tick(time.Second) {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("monitor woring")
			}
		}
	}()

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	select {
	case <-time.After(4 * time.Second):
		fmt.Println("overslept")
	case <-ctx1.Done():
		fmt.Println("===",ctx1.Err())
	}


	ctx2 := context.WithValue(context.Background(), ctxKey("a"), "a")
	ctx2 = context.WithValue(ctx2, "test", "【监控1】")
	get := func(ctx context.Context, k ctxKey) {
		if v, ok := ctx.Value(k).(string); ok {
			fmt.Println(v)
		}
	}
	get(ctx2, ctxKey("a"))
	get(ctx2, ctxKey("b"))
	fmt.Println(ctx2.Value("test"))

}
